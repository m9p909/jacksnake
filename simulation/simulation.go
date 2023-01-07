package simulation

import (
	. "jacksnake/models"
)

func getSnakes(state GameState) []Battlesnake {
	numSnakes := len(state.Board.Snakes) + 1
	snakes := make([]Battlesnake, numSnakes)
	for index, snek := range state.Board.Snakes {
		snakes[index] = snek
	}
	snakes[numSnakes-1] = state.You
	return snakes
}

func removeEndOfTail(snek Battlesnake) Battlesnake {
	snek.Body = snek.Body[:len(snek.Body)-1]
	return snek
}

func ApplyMove(coord Coord, move string) Coord {

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

func getNewHead(head Coord, move string) Coord {
	return ApplyMove(head, move)
}

func addHeadToFront(snek Battlesnake, newHead Coord) Battlesnake {
	newBody := []Coord{
		newHead,
	}
	snek.Body = append(newBody, snek.Body...)
	return snek
}

func updateSnake(snek Battlesnake, move string) Battlesnake {
	snek = removeEndOfTail(snek)
	snek.Head = ApplyMove(snek.Head, move)
	snek = addHeadToFront(snek, snek.Head)
	return snek
}

// assume move is valid
func SimulateMove(state GameState, move string) GameState {
	state.You = updateSnake(state.You, move)

	return state
}

func findSnake(snake Battlesnake, state *GameState) *Battlesnake {
	if snake.ID == state.You.ID {
		return &state.You
	}
	for _, boardSnake := range state.Board.Snakes {
		if snake.ID == boardSnake.ID {
			return &boardSnake
		}
	}
	return nil
}

func SimulateMoveBySnake(state GameState, move string, snake Battlesnake) GameState {
	snekRef := findSnake(snake, &state)

	*snekRef = updateSnake(*snekRef, move)

	return state

}
