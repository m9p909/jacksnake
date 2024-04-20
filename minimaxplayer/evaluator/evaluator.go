package evaluator

import (
	"fmt"
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
	res := math.Pow(float64(snake.Health)/100, 2)
	if res > 1 {
		fmt.Println("Snake", snake)
		panic(res)
	}
	return res
}

const (
	NONE Elem = iota
	FOOD
	HAZARD
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

func inRange(start uint8, end uint8, value uint8) bool {
	return value >= start && value < end
}

func buildSnakeBoard(snakes []Snake, height uint8, width uint8) [][]uint8 {
	board := makeEmptyBoard(height, width)

	for _, snake := range snakes {
		if snake.Health > 0 {
			for _, body := range snake.Body {
				board[body.Y][body.X] = uint8(snake.ID)
			}
		}
	}

	return board
}

func countAvailableSquares(snakes [][]uint8, head Point, b *GameBoard) int {
	q := []Point{head}
	halfTheBoard := int(b.Height * b.Width / 2)

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
		if size >= halfTheBoard {
			return halfTheBoard
		}
	}
	return size
}

func evaluateSpaceConstraint(state *GameBoard, snakeId SnakeID) float64 {
	snake := findSnakeById(&state.Snakes, snakeId)
	snakes := state.Snakes
	snakesBoard := buildSnakeBoard(snakes, state.Height, state.Width)
	availableSquares := countAvailableSquares(snakesBoard, snake.Body[0], state)
	return math.Pow(float64(availableSquares)/float64(state.Height*state.Width/2+1), 1.5)
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
	count := 0.0
	if len(board.Snakes) == 1 {
		return 1
	}

	for i, snake := range board.Snakes {
		if board.Snakes[i].ID != targetSnake.ID {
			score = score + getHealthScore(&snake) // between 0 and 1
			count++
		}
	}
	numSnakes := float64(len(board.Snakes) - 1)
	final := 1 - score/count
	if final < 0 {
		fmt.Println(board)
		fmt.Println("Snakes: ", numSnakes)
		fmt.Println("score: ", score)
		panic(final)
	}
	return final
}

func evaluateDeathScore(snake *Snake) float64 {
	if snake.Health == 0 {
		return 0
	} else {
		return 1
	}
}

func (*SimpleEvaluator) EvaluateBoard(board *GameBoard, snakeId SnakeID, complete bool, count int) float64 {
	snake := findSnakeById(&board.Snakes, snakeId)
	if snake != nil {
		healthScore := getHealthScore(snake)
		if healthScore < 0 || healthScore > 1 {
			fmt.Println(healthScore)
			panic("bad health score")
		}
		otherSnakesHealth := getOtherSnakeHealthScore(board, snake)
		// spaceScore := evaluateSpaceConstraint(board, snakeId)
		// deathScore := evaluateDeathScore(snake)
		score := healthScore*0.8 + otherSnakesHealth*0.2
		// if the max depth is reached
		if score < 0 {
			println("neg score")
		}

		if score > 1 {
			println("score too big")
		}
		// reduce weight of score if not at end of game
		if !complete {
			score = score * 0.01 * (float64(count) + 1)
		}
		if score <= 0 || score > 1 {
			panic("Invalid score")
		}
		return score
	}
	// println("no snake found, this should never happen")
	return 0
}
