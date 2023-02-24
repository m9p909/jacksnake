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

func findSnakeById(snakes *[]Snake, id string) *Snake {
	for _, snake := range *snakes {
		if snake.ID == id {
			return &snake
		}
	}
	return nil
}

func getHealthScore(snake *Snake) float64 {
	target := 80.0
	value := float64(snake.Health)
	difference := math.Abs(target-value) / target
	return math.Abs(1 - difference)
}

func makeEmptyBoard(height int, width int) [][]string {
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

func buildSnakeBoard(snakes []Snake, height int, width int) [][]string {
	board := makeEmptyBoard(height, width)

	for i, snake := range snakes {
		for _, body := range snake.Body {
			board[body.Y][body.X] = strconv.Itoa(i)
		}
	}

	return board
}

func countAvailableSquares(snakes [][]string, head Point) int {
	width := len(snakes[0])
	height := len(snakes[1])
	data := make([][]int, len(snakes))
	for i := range data {
		row := make([]int, len(snakes[0]))
		for j := range row {
			row[j] = 2000
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
				data[front.Y][front.X] != 2000 ||
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

func evaluateSpaceConstraint(state *GameBoard, snakeId string) float64 {
	snake := findSnakeById(&state.Snakes, snakeId)
	snakes := state.Snakes
	snakesBoard := buildSnakeBoard(snakes, state.Height, state.Width)
	availableSquares := countAvailableSquares(snakesBoard, snake.Body[0])
	return math.Pow(float64(availableSquares)/float64(state.Height*state.Width), 1.5)
}

func evaluateDeadSnakes(state *GameBoard, snakeId string) float64 {
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

func (*SimpleEvaluator) EvaluateBoard(board *GameBoard, snakeId string) float64 {
	snake := findSnakeById(&board.Snakes, snakeId)
	if snake != nil {
		healthScore := getHealthScore(snake)
		// spaceScore := evaluateSpaceConstraint(board, snakeId)
		return healthScore
	}
	println("no snake found, this should never happen")
	return 0
}
