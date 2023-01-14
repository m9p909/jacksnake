package safemoves

import (
	"errors"

	"github.com/BattlesnakeOfficial/rules"
)

func ApplyMove(coord rules.Point, move string) rules.Point {

	if move == "up" {
		coord.Y += 1
	}

	if move == "down" {
		coord.Y -= 1
	}

	if move == "left" {
		coord.X -= 1
	}

	if move == "right" {
		coord.X += 1
	}

	return coord
}

func Equals(coord1 rules.Point, coord2 rules.Point) bool {
	return coord1.X == coord2.X && coord1.Y == coord2.Y
}

func getSnake(board rules.BoardState, snakeID string) (*rules.Snake, error) {
	for _, snake := range board.Snakes {
		if snake.ID == snakeID {
			return &snake, nil
		}
	}
	return nil, errors.New("cant find snake")
}

func GetSafeMovesBySnake(state rules.BoardState, snakeID string) []string {
	snake, _ := getSnake(state, snakeID)
	isMoveSafe := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := snake.Body[0] // Coordinates of your head
	myNeck := snake.Body[1] // Coordinates of your "neck"

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
	boardWidth := state.Width
	boardHeight := state.Height

	if snake.Body[0].Y == boardHeight-1 {
		isMoveSafe["up"] = false
	}

	if snake.Body[0].Y == 0 {
		isMoveSafe["down"] = false
	}

	if snake.Body[0].X == 0 {
		isMoveSafe["left"] = false
	}

	if snake.Body[0].X == boardWidth-1 {
		isMoveSafe["right"] = false
	}

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	// mybody := state.You.Body
	mybody := snake.Body
	for move, isSafe := range isMoveSafe {
		if isSafe {
			nextHead := ApplyMove(myHead, move)
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
	opponents := state.Snakes

	for move, isSafe := range isMoveSafe {
		if isSafe {
			next_head := ApplyMove(myHead, move)
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

	// hazards
	for move, isSafe := range isMoveSafe {
		if isSafe {
			next_head := ApplyMove(myHead, move)
			for _, hazard := range state.Hazards {
				if Equals(hazard, next_head) {
					isMoveSafe[move] = false
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
