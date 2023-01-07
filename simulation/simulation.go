package simulation

import (
	. "jacksnake/models"
)

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

func updateSnakePosition(snek Battlesnake, move string) Battlesnake {
	snek = removeEndOfTail(snek)
	snek.Head = ApplyMove(snek.Head, move)
	snek = addHeadToFront(snek, snek.Head)
	return snek
}

func updateSnakeFood(snek Battlesnake, state GameState) GameState {
	for _, food := range state.Board.Food {
		if Equals(snek.Head, food) {
			state.Board.Food

		}
	}
	return state
}

func findSnake(snake Battlesnake, state *GameState) *Battlesnake {
	for _, boardSnake := range state.Board.Snakes {
		if snake.ID == boardSnake.ID {
			return &boardSnake
		}
	}
	print("could not find snek in simulation")
	return nil
}

func SimulateMoveBySnake(state GameState, move string, snake Battlesnake) GameState {
	snekRef := findSnake(snake, &state)

	*snekRef = updateSnakePosition(*snekRef, move)

	return state

}
