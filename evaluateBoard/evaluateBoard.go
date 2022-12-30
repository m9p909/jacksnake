package evaluateboard

import (
	"fmt"
	. "jacksnake/models"
	"math"
	"strconv"
)

type GameData struct {
	SnakeBoard [][]string // snake
	Snakes     []Battlesnake
	FoodBoard  [][]string
	SnakeID    int
	State      GameState
}

func getSnakes(state GameState) []Battlesnake {
	numSnakes := len(state.Board.Snakes) + 1
	snakes := make([]Battlesnake, numSnakes)
	for index, snek := range state.Board.Snakes {
		snakes[index] = snek
	}
	snakes[numSnakes-1] = state.You
	return snakes
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

func buildSnakeBoard(snakes []Battlesnake, height int, width int) [][]string {
	board := makeEmptyBoard(height, width)

	for i, snake := range snakes {
		for _, body := range snake.Body {
			board[body.Y][body.X] = strconv.Itoa(i)
		}
	}

	return board

}

func displayBoard(board [][]string) {
	for _, row := range board {
		for _, col := range row {
			print(col)
		}
		print("\n")
	}
}

func getFood(food []Coord, height int, width int) [][]string {
	board := makeEmptyBoard(height, width)

	for _, foods := range food {

		board[foods.Y][foods.X] = "x"

	}

	return board

}

func buildGameData(state GameState) GameData {
	snakes := getSnakes(state)

	return GameData{
		SnakeBoard: buildSnakeBoard(snakes, state.Board.Height, state.Board.Width),
		Snakes:     snakes,
		FoodBoard:  getFood(state.Board.Food, state.Board.Height, state.Board.Width),
		SnakeID:    len(snakes) - 1,
	}

}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func createDistanceGraph(snakes [][]string, snake Battlesnake) [][]int {
	width := len(snakes[0])
	height := len(snakes[1])
	var data = make([][]int, len(snakes))
	for i := range data {
		row := make([]int, len(snakes[0]))
		for j := range row {
			row[j] = 2000
		}
		data[i] = row
	}

	q := []Coord{snake.Head}

	dist := 0
	for len(q) > 0 {
		nextCoords := []Coord{}

		for _, front := range q {
			if front.X >= width ||
				front.X < 0 ||
				front.Y >= height ||
				front.Y < 0 ||
				data[front.Y][front.X] != 2000 ||
				(snakes[front.Y][front.X] != "-" && !Equals(front, snake.Head)) {
				continue
			}
			data[front.Y][front.X] = min(data[front.Y][front.X], dist)
			nextQ := []Coord{
				{X: front.X + 1, Y: front.Y},
				{X: front.X - 1, Y: front.Y},
				{X: front.X, Y: front.Y + 1},
				{X: front.X, Y: front.Y - 1},
			}
			nextCoords = append(nextCoords, nextQ...)
		}

		q = nextCoords

		dist++
	}

	for i, row := range data {
		for j, col := range row {
			if col == 2000 {
				data[i][j] = -1
			}
		}
	}

	return data

}

func getFoodScore(food []Coord, distanceGraph [][]int) float64 {
	score := 0.0
	width := float64(len(distanceGraph[0]))
	height := float64(len(distanceGraph))

	for _, f := range food {
		if distanceGraph[f.Y][f.X] != -1 {
			dist := float64(distanceGraph[f.Y][f.X])
			score += (math.Pow(((width + height) - dist), 2)) // my best snake is the one that avoids food?
		}
	}

	return math.Tanh(score / 1000)
}

func EvaluateFoodScore(state GameState) float64 {
	snakes := getSnakes(state)
	snakesBoard := buildSnakeBoard(snakes, state.Board.Height, state.Board.Width)
	distanceGraph := createDistanceGraph(snakesBoard, state.You)
	score := getFoodScore(state.Board.Food, distanceGraph)
	health := float64(state.You.Health) / 100
	if health < 0.25 {
		return score * (1 - health)
	} else {
		return 1 - score
	}
}

func countAvailableSquares(snakes [][]string, head Coord) int {
	width := len(snakes[0])
	height := len(snakes[1])
	var data = make([][]int, len(snakes))
	for i := range data {
		row := make([]int, len(snakes[0]))
		for j := range row {
			row[j] = 2000
		}
		data[i] = row
	}

	q := []Coord{head}

	size := 0
	for len(q) > 0 {
		nextCoords := []Coord{}

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
			nextQ := []Coord{
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

func EvaluateSpaceConstraint(state GameState) float64 {
	snakes := getSnakes(state)
	snakesBoard := buildSnakeBoard(snakes, state.Board.Height, state.Board.Width)
	availableSquares := countAvailableSquares(snakesBoard, state.You.Head)
	return math.Pow(float64(availableSquares)/float64(state.Board.Height*state.Board.Width), 1.5)
}

func evaluateBundling(state GameState) float64 {
	snakes := getSnakes(state)
	snakesBoard := buildSnakeBoard(snakes, state.Board.Height, state.Board.Width)
	you := state.You
	width := state.Board.Width
	height := state.Board.Height
	count := 0
	for _, j := range []int{1, -1} {
		newhead := Coord{X: you.Head.X + j, Y: you.Head.Y}
		if newhead.X >= width ||
			newhead.X < 0 {
			continue
		}

		if snakesBoard[newhead.Y][newhead.X] == fmt.Sprintf("%d", len(snakes)-1) {
			count++
		}
	}

	for _, j := range []int{1, -1} {
		newhead := Coord{X: you.Head.X, Y: you.Head.Y + j}
		if newhead.Y >= height ||
			newhead.Y < 0 {
			continue
		}

		if snakesBoard[newhead.Y][newhead.X] == fmt.Sprintf("%d", len(snakes)-1) {
			count++
		}
	}

	return float64(count) / 4
}

func printScore(turn int, score float64) {

	println("Turn: ", turn, " Scored res: ", score)
}

func printWeigths(foodScore float64, space float64, bundling float64) {
	println("Food: ", foodScore)
	println("Space: ", space)
	println("Bundling: ", bundling)
}

func EvaluateState(state GameState) float64 {
	foodScore := EvaluateFoodScore(state)
	space := EvaluateSpaceConstraint(state)
	bundling := evaluateBundling(state)
	res := foodScore*0.2 + space*0.75 + bundling*0.05
	return res
}
