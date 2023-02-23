package minimaxplayer_test

import (
	"jacksnake/minimaxplayer"
	. "jacksnake/models"
	"testing"
)

func getGameStateTest2() GameState {
	/*
		4 - - - -
		3 - 0 - f
		2 0 0 1 1
		1 x - 1 f
		0 - x - -
			0 1 2 3
	*/
	snakes := []Battlesnake{
		{
			ID:     "1",
			Health: 99,
			Head:   Coord{X: 0, Y: 2},
			Body:   []Coord{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 1, Y: 3}},
		},
		{
			ID:     "2",
			Health: 100,
			Head:   Coord{X: 3, Y: 2},
			Body:   []Coord{{X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}},
		},
	}

	food := []Coord{
		{X: 3, Y: 1},
		{X: 3, Y: 3},
	}

	hazards := []Coord{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	state := GameState{
		Turn: 2,
		You:  snakes[0],
		Board: Board{
			Snakes: []Battlesnake{
				snakes[0], snakes[1],
			},
			Food:    food,
			Hazards: hazards,
			Width:   4,
			Height:  4,
		},
	}

	return state
}

func randomPlayerTest(t *testing.T, state GameState, badMoves []string) {
	player := minimaxplayer.BuildRandomPlayer()
	move := player.Move(state)
	for _, badmove := range badMoves {
		if move == badmove {
			println("cannot go ", move)
			t.FailNow()
		}
	}
}

func Test_randomPlayer(t *testing.T) {
	randomPlayerTest(t, getGameStateTest1(), []string{"right", "up"})
	randomPlayerTest(t, getGameStateTest2(), []string{"left", "down", "right"})
}
