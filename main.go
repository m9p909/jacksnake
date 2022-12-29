package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"fmt"
	"jacksnake/evaluateBoard"
	. "jacksnake/models"
	"log"
	"math/rand"
	"time"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "m9p909",  // TODO: Your Battlesnake username
		Color:      "#b13859", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

func equal(coord1 Coord, coord2 Coord) bool {
	return coord1.X == coord2.X && coord1.Y == coord2.Y
}

func GetSafeMoves(state GameState) []string {
	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		isMoveSafe["left"] = false

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		isMoveSafe["right"] = false

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		isMoveSafe["down"] = false

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		isMoveSafe["up"] = false
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	if state.You.Body[0].Y == boardHeight-1 {
		isMoveSafe["up"] = false
	}

	if state.You.Body[0].Y == 0 {
		println("cannot go down")
		isMoveSafe["down"] = false
	}

	if state.You.Body[0].X == 0 {
		isMoveSafe["left"] = false
	}

	if state.You.Body[0].X == boardWidth-1 {
		isMoveSafe["right"] = false
	}

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	// mybody := state.You.Body
	mybody := state.You.Body
	for move, isSafe := range isMoveSafe {
		if isSafe {
			nextHead := ApplyMove(myHead, move)
			for index, coord := range mybody {
				if index != 0 {
					if equal(nextHead, coord) {
						isMoveSafe[move] = false
					}
				}

			}
		}
	}

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	opponents := state.Board.Snakes

	for move, isSafe := range isMoveSafe {
		if isSafe {
			next_head := ApplyMove(myHead, move)
			for _, snake := range opponents {
				for _, body := range snake.Body {
					if next_head == body {
						if equal(next_head, body) {
							isMoveSafe[move] = false
						}
					}
				}
			}
		}
	}

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	return safeMoves
}

func determineBestMove(state GameState, safeMoves []string) string {
	if len(safeMoves) <= 0 {
		println("no safe moves")
		return "down"
	}
	max := 0.0
	maxMove := ""
	for _, move := range safeMoves {
		newState := simulateMove(state, move)
		val := evaluateboard.EvaluateState(newState)
		if val > max {
			max = val
			maxMove = move
		}
	}

	if maxMove == "" {
		println("could not determine move, no good moves, picking randomly")
		return determineRandomMove(safeMoves)

	}
	return maxMove
}

func determineRandomMove(safeMoves []string) string {
	return safeMoves[rand.Intn(len(safeMoves))]
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {
	t1 := time.Now()

	safeMoves := GetSafeMoves(state)

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := determineBestMove(state, safeMoves)

	fmt.Printf("time: %s\n", time.Since(t1))

	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
