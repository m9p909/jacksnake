package models

import "strconv"

const (
	hazardSymbol = "x"
	foodSymbol   = "f"
)

type Grid struct {
	Board         [][]string
	InitialSnakes []Battlesnake
	HazardSymbol  string
	FoodSymbol    string
}

func NewGrid(state GameState) Grid {
	grid := Grid{
		Board:         make([][]string, state.Board.Height),
		InitialSnakes: state.Board.Snakes,
		HazardSymbol:  hazardSymbol,
		FoodSymbol:    foodSymbol,
	}

	for i := range grid.Board {
		grid.Board[i] = make([]string, state.Board.Width)
	}

	for i, snake := range grid.InitialSnakes {
		for _, body := range snake.Body {
			grid.Board[body.Y][body.X] = strconv.Itoa(i)
		}
	}

	for _, food := range state.Board.Food {
		grid.Board[food.Y][food.X] = grid.FoodSymbol
	}

	for _, hazard := range state.Board.Hazards {
		grid.Board[hazard.Y][hazard.X] = grid.HazardSymbol
	}

	return grid
}
