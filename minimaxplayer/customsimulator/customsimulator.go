package customsimulator

import (
	"errors"
	"fmt"
	. "jacksnake/minimaxplayer/coreplayer"
	"sort"
)

type CustomSimulator struct{}

func (*CustomSimulator) SimulateMoves(board *GameBoard, moves []SnakeMove) {
	_, err := MoveSnakesStandard(board, moves)
	if err != nil {
		fmt.Println(err)
	}
	_, err = ReduceSnakeHealthStandard(board, moves)
	if err != nil {
		fmt.Println(err)
	}
	_, err = DamageHazardsStandard(board, moves)
	if err != nil {
		fmt.Println(err)
	}
	EliminateSnakesStandard(board, moves)
}

func (*CustomSimulator) GetValidMoves(board *GameBoard, snakeId SnakeID) []Direction {
	return []Direction{UP, DOWN, LEFT, RIGHT}
	// broken
	// return GetSafeMovesBySnake(board, snakeId)
}

func New() Simulator {
	return &CustomSimulator{}
}

// ValidMoves

func ApplyMove(coord Point, move Direction) Point {
	if move == UP {
		coord.Y += 1
	}

	if move == DOWN {
		coord.Y -= 1
	}

	if move == LEFT {
		coord.X -= 1
	}

	if move == RIGHT {
		coord.X += 1
	}

	return coord
}

func getSnake(board *GameBoard, snakeID SnakeID) (*Snake, error) {
	for _, snake := range board.Snakes {
		if snake.ID == snakeID {
			return &snake, nil
		}
	}
	return nil, errors.New("cant find snake")
}

func GetSafeMovesBySnake(state *GameBoard, snakeID SnakeID) []Direction {
	snake, _ := getSnake(state, snakeID)
	isMoveSafe := map[Direction]bool{
		UP:    true,
		DOWN:  true,
		LEFT:  true,
		RIGHT: true,
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	boardWidth := state.Width
	boardHeight := state.Height

	if snake.Body[0].Y == boardHeight-1 {
		isMoveSafe[UP] = false
	}

	if snake.Body[0].Y == 0 {
		isMoveSafe[DOWN] = false
	}

	if snake.Body[0].X == 0 {
		isMoveSafe[LEFT] = false
	}

	if snake.Body[0].X == boardWidth-1 {
		isMoveSafe[RIGHT] = false
	}

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	// mybody := state.You.Body
	mybody := snake.Body
	for move, isSafe := range isMoveSafe {
		if isSafe {
			nextHead := ApplyMove(snake.Body[0], move)
			var body []Point
			// if snake health is 100, does that mean it just ate
			if snake.Health == 100 {
				body = mybody
			} else {
				body = mybody[:len(mybody)-1]
			}
			for index, coord := range body {
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
			next_head := ApplyMove(snake.Body[0], move)
			for _, snake := range opponents {
				var body []Point
				if snake.Health == 100 {
					body = mybody
				} else {
					body = mybody[:len(mybody)-1]
				}
				for _, body := range body {
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
			next_head := ApplyMove(snake.Body[0], move)
			for _, hazard := range state.Hazards {
				if Equals(hazard, next_head) {
					isMoveSafe[move] = false
				}
			}
		}
	}

	// Are there any safe moves left?
	safeMoves := []Direction{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}
	return safeMoves
}

// Simulator

// copied from official rules

func MoveSnakesStandard(b *GameBoard, moves []SnakeMove) (bool, error) {
	// no-op when moves are empty
	if len(moves) == 0 {
		return false, nil
	}

	// Sanity check that all non-eliminated snakes have moves and bodies.
	for i := 0; i < len(b.Snakes); i++ {
		snake := &b.Snakes[i]
		if snake.Health != 0 {
			continue
		}

		if len(snake.Body) == 0 {
			return false, errors.New("snake has zero length")
		}

		moveFound := false
		for _, move := range moves {
			if snake.ID == move.ID {
				moveFound = true
				break
			}
		}
		if !moveFound {
			return false, errors.New("no move found")
		}
	}

	for i := 0; i < len(b.Snakes); i++ {
		snake := &b.Snakes[i]
		if snake.Health == 0 {
			continue
		}

		for _, move := range moves {
			if move.ID == snake.ID {
				appliedMove := move.Move
				switch move.Move {
				case UP, DOWN, LEFT, RIGHT:
					break
				default:
					appliedMove = getDefaultMove(snake.Body)
				}

				newHead := Point{}
				switch appliedMove {
				// Guaranteed to be one of these options given the clause above
				case UP:
					newHead.X = snake.Body[0].X
					newHead.Y = snake.Body[0].Y + 1
				case DOWN:
					newHead.X = snake.Body[0].X
					newHead.Y = snake.Body[0].Y - 1
				case LEFT:
					newHead.X = snake.Body[0].X - 1
					newHead.Y = snake.Body[0].Y
				case RIGHT:
					newHead.X = snake.Body[0].X + 1
					newHead.Y = snake.Body[0].Y
				}

				if len(snake.Body) == 2 && Equals(snake.Body[1], newHead) {
					snake.Health = 0
				}

				// Append new head, pop old tail
				// if health is 100 assume just ate, no reduction in length
				if snake.Health == 100 {
					snake.Body = append([]Point{newHead}, snake.Body...)
				} else {
					snake.Body = append([]Point{newHead}, snake.Body[:len(snake.Body)-1]...)
				}
				// handle edge case when snake body has length 2
			}
		}
	}
	return false, nil
}

func getDefaultMove(snakeBody []Point) Direction {
	if len(snakeBody) >= 2 {
		// Use neck to determine last move made
		head, neck := snakeBody[0], snakeBody[1]
		// Situations where neck is next to head
		if head.X == neck.X+1 {
			return RIGHT
		} else if head.X == neck.X-1 {
			return LEFT
		} else if head.Y == neck.Y+1 {
			return UP
		} else if head.Y == neck.Y-1 {
			return DOWN
		}
		// Consider the wrapped cases using zero axis to anchor
		if head.X == 0 && neck.X > 0 {
			return RIGHT
		} else if neck.X == 0 && head.X > 0 {
			return LEFT
		} else if head.Y == 0 && neck.Y > 0 {
			return UP
		} else if neck.Y == 0 && head.Y > 0 {
			return DOWN
		}
	}
	return UP
}

func ReduceSnakeHealthStandard(b *GameBoard, _ []SnakeMove) (bool, error) {
	for i := 0; i < len(b.Snakes); i++ {
		if b.Snakes[i].Health > 0 {
			b.Snakes[i].Health = b.Snakes[i].Health - 1
		}
	}
	return false, nil
}

func DamageHazardsStandard(b *GameBoard, _ []SnakeMove) (bool, error) {
	var hazardDamage uint8 = 1
	var SnakeMaxHealth uint8 = 100
	for i := 0; i < len(b.Snakes); i++ {
		snake := &b.Snakes[i]
		if snake.Health == 0 {
			continue
		}
		head := snake.Body[0]
		for _, p := range b.Hazards {
			if head == p {
				// If there's a food in this square, don't reduce health
				foundFood := false
				for _, food := range b.Food {
					if p == food {
						foundFood = true
					}
				}
				if foundFood {
					continue
				}

				// Snake is in a hazard, reduce health
				snake.Health = snake.Health - hazardDamage
				if snake.Health < 0 {
					snake.Health = 0
				}
				if snake.Health > SnakeMaxHealth {
					snake.Health = SnakeMaxHealth
				}
			}
		}
	}

	return false, nil
}

func EliminateSnakesStandard(b *GameBoard, _ []SnakeMove) (bool, error) {
	// First order snake indices by length.
	// In multi-collision scenarios we want to always attribute elimination to the longest snake.
	snakeIndicesByLength := make([]int, len(b.Snakes))
	for i := 0; i < len(b.Snakes); i++ {
		snakeIndicesByLength[i] = i
	}
	sort.Slice(snakeIndicesByLength, func(i int, j int) bool {
		lenI := len(b.Snakes[snakeIndicesByLength[i]].Body)
		lenJ := len(b.Snakes[snakeIndicesByLength[j]].Body)
		return lenI > lenJ
	})

	// First, iterate over all non-eliminated snakes and eliminate the ones
	// that are out of health or have moved out of bounds.
	for i := 0; i < len(b.Snakes); i++ {
		snake := &b.Snakes[i]
		if snake.Health == 0 {
			continue
		}
		if len(snake.Body) <= 0 {
			return false, errors.New("snake 0 length body")
		}

		if snakeIsOutOfHealth(snake) {
			snake.Health = 0
			continue
		}

		if snakeIsOutOfBounds(snake, b.Width, b.Height) {
			snake.Health = 0
			continue
		}
	}

	// Next, look for any collisions. Note we apply collision eliminations
	// after this check so that snakes can collide with each other and be properly eliminated.

	collisionEliminations := []SnakeID{}
	for i := 0; i < len(b.Snakes); i++ {
		snake := &b.Snakes[i]
		if snake.Health == 0 {
			continue
		}
		if len(snake.Body) <= 0 {
			return false, errors.New("snake 0 length body")
		}

		// Check for self-collisions first
		if snakeHasBodyCollided(snake, snake) {
			collisionEliminations = append(collisionEliminations, snake.ID)
			continue
		}

		// Check for body collisions with other snakes second

		for _, otherIndex := range snakeIndicesByLength {
			other := &b.Snakes[otherIndex]
			if other.Health == 0 {
				continue
			}
			if snake.ID != other.ID && snakeHasBodyCollided(snake, other) {
				collisionEliminations = append(collisionEliminations, snake.ID)
				break
			}
		}

		// Check for head-to-heads last
		for _, otherIndex := range snakeIndicesByLength {
			other := &b.Snakes[otherIndex]
			if other.Health == 0 {
				continue
			}
			if snake.ID != other.ID && snakeHasLostHeadToHead(snake, other) {
				collisionEliminations = append(collisionEliminations, snake.ID)
				break
			}
		}
	}

	// kill the snakes
	for i := range b.Snakes {
		for j := range collisionEliminations {
			if b.Snakes[i].ID == collisionEliminations[j] {
				b.Snakes[i].Health = 0
			}
		}
	}
	return false, nil
}

func snakeIsOutOfHealth(s *Snake) bool {
	return s.Health <= 0
}

func snakeIsOutOfBounds(s *Snake, boardWidth uint8, boardHeight uint8) bool {
	for _, point := range s.Body {
		if (point.X < 0) || (point.X >= boardWidth) {
			return true
		}
		if (point.Y < 0) || (point.Y >= boardHeight) {
			return true
		}
	}
	return false
}

func snakeHasBodyCollided(s *Snake, other *Snake) bool {
	head := s.Body[0]
	for i, body := range other.Body {
		if i == 0 {
			continue
		} else if head.X == body.X && head.Y == body.Y {
			return true
		}
	}
	return false
}

func snakeHasLostHeadToHead(s *Snake, other *Snake) bool {
	if s.Body[0].X == other.Body[0].X && s.Body[0].Y == other.Body[0].Y {
		return len(s.Body) <= len(other.Body)
	}
	return false
}

func FeedSnakesStandard(b *GameBoard, _ []SnakeMove) (bool, error) {
	newFood := []Point{}
	for _, food := range b.Food {
		foodHasBeenEaten := false
		for i := 0; i < len(b.Snakes); i++ {
			snake := &b.Snakes[i]

			// Ignore eliminated and zero-length snakes, they can't eat.
			if snake.Health == 0 || len(snake.Body) == 0 {
				continue
			}

			if snake.Body[0].X == food.X && snake.Body[0].Y == food.Y {
				feedSnake(snake)
				foodHasBeenEaten = true
			}
		}
		// Persist food to next BoardState if not eaten
		if !foodHasBeenEaten {
			newFood = append(newFood, food)
		}
	}

	b.Food = newFood
	return false, nil
}

func feedSnake(snake *Snake) {
	growSnake(snake)
	snake.Health = 100
}

func growSnake(snake *Snake) {
	if len(snake.Body) > 0 {
		snake.Body = append(snake.Body, snake.Body[len(snake.Body)-1])
	}
}

func GameOverStandard(b *GameBoard, moves []SnakeMove) (bool, error) {
	numSnakesRemaining := 0
	for i := 0; i < len(b.Snakes); i++ {
		if b.Snakes[i].Health > 0 {
			numSnakesRemaining++
		}
	}
	return numSnakesRemaining <= 1, nil
}
