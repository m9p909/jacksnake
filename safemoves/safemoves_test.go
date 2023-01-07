package safemoves

import (
	. "jacksnake/models"
	"testing"
)

func Test_getSafeMovesBySnake(t *testing.T) {
	/*
		3 - 0 - 1
		2 - 0 1 1
		1 - - 1 -
		0 - - - -
			0 1 2 3
	*/
	snakes := []Battlesnake{
		{Head: Coord{X: 1, Y: 2},
			Body: []Coord{{X: 1, Y: 2}, {X: 1, Y: 3}}},
		{Head: Coord{X: 3, Y: 3},
			Body: []Coord{{X: 3, Y: 3}, {X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}}}}

	state := GameState{
		You: snakes[0],
		Board: Board{
			Snakes: []Battlesnake{
				snakes[1],
			},
			Width:  4,
			Height: 4,
		},
	}

	res := GetSafeMovesBySnake(state, snakes[1])
	for i := range res {
		print(res[i], " ")
	}
	println()
	if len(res) != 1 {
		println("snake 1 should only have 1 option")
		t.FailNow()
	}

}
