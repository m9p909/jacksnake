package officialrulesapi_test

import (
	"fmt"
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/officialrulesapi"
	"testing"
)

func getCoreState1() coreplayer.GameBoard {
	/*
		3 - 0 - f
		2 - 0 1 1
		1 x - 1 f
		0 - x - -
			0 1 2 3
	*/

	food := []coreplayer.Point{
		{X: 3, Y: 1},
		{X: 3, Y: 3},
	}

	hazards := []coreplayer.Point{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	snakes := []coreplayer.Snake{
		{
			ID:     "1",
			Body:   []coreplayer.Point{{X: 1, Y: 2}, {X: 1, Y: 3}},
			Health: 99,
		},
		{
			ID:     "2",
			Body:   []coreplayer.Point{{X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}},
			Health: 100,
		},
	}

	return coreplayer.GameBoard{
		Turn:    2,
		Height:  5,
		Width:   4,
		Food:    food,
		Hazards: hazards,
		Snakes:  snakes,
	}
}

func parametrized(t *testing.T, state coreplayer.GameBoard, id string, validMoves map[string]int) {
	rules := officialrulesapi.GetOfficialRules()
	moves := rules.GetValidMoves(state, id)

	if len(moves) != len(validMoves) {
		println("Weird number of moves")
		t.Fail()
	}

	for _, move := range moves {
		_, ok := validMoves[move]
		if !ok {
			fmt.Printf("%+v\n", state)
			println()
			println("bad move ", move)
			fmt.Printf("%+v\n", moves)
			t.Fail()
		}
	}

}

func Test_OfficialRulesAdapter(t *testing.T) {
	parametrized(t, getCoreState1(), "1", map[string]int{"down": 1, "left": 1})
	parametrized(t, getCoreState1(), "2", map[string]int{"up": 1, "down": 1})
}
