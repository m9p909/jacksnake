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
	"jacksnake/minmax"
	. "jacksnake/models"
	safemoves "jacksnake/safemoves"
	simulate "jacksnake/simulation"
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

func determineBestMoveHeuristics(state GameState, safeMoves []string) string {
	if len(safeMoves) <= 0 {
		println("no safe moves")
		return "down"
	}
	max := 0.0
	maxMove := ""
	for _, move := range safeMoves {
		newState := simulate.SimulateMove(state, move)
		val := evaluateboard.EvaluateCurrentState(newState)
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

func determineBestMoveMiniMax(state GameState) string {
	minimaxAlgo := minmax.NewMiniMax(3, state)
	return minimaxAlgo.Analyze()
}

func determineRandomMove(safeMoves []string) string {
	return safeMoves[rand.Intn(len(safeMoves))]
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {
	t1 := time.Now()

	safeMoves := safemoves.GetSafeMoves(state)

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// Choose a random move from the safe ones
	nextMove := determineBestMoveMiniMax(state)

	fmt.Printf("time: %s\n", time.Since(t1))

	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
