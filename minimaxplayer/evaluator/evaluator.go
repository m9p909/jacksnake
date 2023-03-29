package evaluator

import (
	. "jacksnake/minimaxplayer/coreplayer"
	"math"
)

type SimpleEvaluator struct{}

func NewSimpleEvaluator() Evaluator {
	return &SimpleEvaluator{}
}

func findSnakeById(snakes *[]Snake, id SnakeID) *Snake {
	for _, snake := range *snakes {
		if snake.ID == id {
			return &snake
		}
	}
	return nil
}

func getHealthScore(snake *Snake) float64 {
	return float64(snake.Health) / 100
}

const (
	FOOD Elem = iota
	HAZARD
	NONE
)

func makeEmptyBoard(height uint8, width uint8) [][]uint8 {
	board := make([][]uint8, height)
	for b := range board {
		row := make([]uint8, width)
		for val := range row {
			row[val] = uint8(NONE)
		}
		board[b] = row
	}

	return board
}

type Elem uint8

func buildSnakeBoard(snakes []Snake, height uint8, width uint8) [][]uint8 {
	board := makeEmptyBoard(height, width)

	for _, snake := range snakes {
		for _, body := range snake.Body {
			board[body.Y][body.X] = uint8(snake.ID)
		}
	}

	return board
}

func countAvailableSquares(snakes [][]uint8, head Point, b *GameBoard) int {
	q := []Point{head}

	size := 0
	for len(q) > 0 {
		nextCoords := []Point{}

		for _, front := range q {
			if front.X >= b.Width ||
				front.X < 0 ||
				front.Y >= b.Height ||
				front.Y < 0 ||
				snakes[front.Y][front.X] == 250 ||
				(snakes[front.Y][front.X] != uint8(NONE) && !Equals(front, head)) {
				continue
			}
			snakes[front.Y][front.X] = 250
			size++
			nextQ := []Point{
				{X: front.X + 1, Y: front.Y},
				{X: front.X - 1, Y: front.Y},
				{X: front.X, Y: front.Y + 1},
				{X: front.X, Y: front.Y - 1},
			}
			nextCoords = append(nextCoords, nextQ...)
		}
		q = nextCoords
	}
	return size
}

func evaluateSpaceConstraint(state *GameBoard, snakeId SnakeID) float64 {
	snake := findSnakeById(&state.Snakes, snakeId)
	snakes := state.Snakes
	snakesBoard := buildSnakeBoard(snakes, state.Height, state.Width)
	availableSquares := countAvailableSquares(snakesBoard, snake.Body[0], state)
	return math.Pow(float64(availableSquares)/float64(state.Height*state.Width), 1.5)
}

func evaluateDeadSnakes(state *GameBoard, snakeId SnakeID) float64 {
	if len(state.Snakes) > 1 {

		deadSnakes := 0
		for _, snake := range state.Snakes {
			if snake.Health <= 0 && snake.ID != snakeId {
				deadSnakes++
			}
		}
		return float64(deadSnakes) / float64(len(state.Snakes)-1)
	}
	return 0.5
}

func getOtherSnakeHealthScore(board *GameBoard, targetSnake *Snake) float64 {
	score := 0.0
	if len(board.Snakes) == 1 {
		return 1
	}

	for i := range board.Snakes {
		if board.Snakes[i].ID != targetSnake.ID {
			score += getHealthScore(&board.Snakes[i])
		}
	}

	return 1 - (score / (float64(len(board.Snakes) - 1)))
}

func evaluateDeathScore(snake *Snake) float64 {
	if snake.Health == 0 {
		return 0
	} else {
		return 1
	}
}

func (*SimpleEvaluator) EvaluateBoard(board *GameBoard, snakeId SnakeID) float64 {
	snake := findSnakeById(&board.Snakes, snakeId)
	if snake != nil {
		healthScore := getHealthScore(snake)
		otherSnakesHealth := getOtherSnakeHealthScore(board, snake)
		// spaceScore := evaluateSpaceConstraint(board, snakeId)
		// deathScore := evaluateDeathScore(snake)
		return healthScore*0.8 + otherSnakesHealth*0.2
	}
	println("no snake found, this should never happen")
	return 0
}
