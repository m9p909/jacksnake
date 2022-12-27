package main

import "strconv"

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
		board[b] = make([]string, width)
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

func getFood(food []Coord, height int, width int) [][]string {
	board := makeEmptyBoard(height, width)

	for _, foods := range food {

		board[foods.Y][foods.X] = "x"

	}

	return board

}

func buildFoodBoard(state GameState)

func buildGameData(state GameState) GameData {
	snakes := getSnakes(state)

	return GameData{
		SnakeBoard: buildSnakeBoard(snakes, state.Board.Height, state.Board.Width),
		Snakes:     snakes,
		FoodBoard:  getFood(state.Board.Food, state.Board.Height, state.Board.Width),
		SnakeID:    len(snakes) - 1,
	}

}

func evaluateBoard(data GameData) {
	you := data.Snakes[data.SnakeID]

}
