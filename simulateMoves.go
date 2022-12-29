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

func addHeadToFront(snek Battlesnake, newHead Coord) Battlesnake {
	newBody := []Coord{
		newHead,
	}
	snek.Body = append(newBody, snek.Body...)
	return snek
}

func updateYou(snek Battlesnake, move string) Battlesnake {
	snek = removeEndOfTail(snek)
	snek.Head = ApplyMove(snek.Head, move)
	snek = addHeadToFront(snek, snek.Head)
	return snek
}

// assume move is valid
func simulateMove(state GameState, move string) GameState {
	state.You = updateYou(state.You, move)

	return state
}
