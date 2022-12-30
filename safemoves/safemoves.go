package safemoves

import (
	. "jacksnake/models"
	"jacksnake/simulation"
)

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
			nextHead := simulation.ApplyMove(myHead, move)
			for index, coord := range mybody {
				if index != 0 {
					if Equals(nextHead, coord) {
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
			next_head := simulation.ApplyMove(myHead, move)
			for _, snake := range opponents {
				for _, body := range snake.Body {
					if next_head == body {
						if Equals(next_head, body) {
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
