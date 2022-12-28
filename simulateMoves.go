package main

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

func addHeadToFront(snek: Battlesnake) {

}

// assume move is valid
func simulateMove(state GameState, move string) GameState {
	snek := state.You
	snek = removeEndOfTail(snek)
	
	snek.Head = ApplyMove(snek.Head)
}
