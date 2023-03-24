package evaluator

import (
	. "jacksnake/minimaxplayer/coreplayer"
	"math"
	"strconv"
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

func makeEmptyBoard(height uint8, width uint8) [][]string {
	board := make([][]string, height)
	for b := range board {
		row := make([]string, width)
		for val := range row {
			row[val] = "-"
		}
		board[b] = row
	}

	return board
}

func buildSnakeBoard(snakes []Snake, height uint8, width uint8) [][]string {
	board := makeEmptyBoard(height, width)

	for i, snake := range snakes {
		for _, body := range snake.Body {
			board[body.Y][body.X] = strconv.Itoa(i)
		}
	}

	return board
}

func countAvailableSquares(snakes [][]string, head Point) int {
	width := uint8(len(snakes[0]))
	height := uint8(len(snakes[1]))
	data := make([][]int, len(snakes))
	for i := range data {
		row := make([]int, len(snakes[0]))
		for j := range row {
			row[j] = 250
		}
		data[i] = row
	}

	q := []Point{head}

	size := 0
	for len(q) > 0 {
		nextCoords := []Point{}

		for _, front := range q {
			if front.X >= width ||
				front.X < 0 ||
				front.Y >= height ||
				front.Y < 0 ||
				data[front.Y][front.X] != 250 ||
				(snakes[front.Y][front.X] != "-" && !Equals(front, head)) {
				continue
			}
			data[front.Y][front.X] = 1
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
	availableSquares := countAvailableSquares(snakesBoard, snake.Body[0])
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

	return 1 - (score / float64(len(board.Snakes)-1))
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
		// healthScore := getHealthScore(snake)
		// otherSnakesHealth := getOtherSnakeHealthScore(board, snake)
		// spaceScore := evaluateSpaceConstraint(board, snakeId)
		deathScore := evaluateDeathScore(snake)
		return deathScore
	}
	println("no snake found, this should never happen")
	return 0
}
