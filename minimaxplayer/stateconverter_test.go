package minimaxplayer_test

import (
	"bytes"
	"encoding/json"
	"jacksnake/minimaxplayer"
	"jacksnake/minimaxplayer/coreplayer"
	. "jacksnake/models"
	"testing"
)

func getGameStateTest1() GameState {
	/*
		3 - 0 - f
		2 - 0 1 1
		1 x - 1 f
		0 - x - -
			0 1 2 3
	*/
	snakes := []Battlesnake{
		{
			ID:     "1",
			Health: 99,
			Head:   Coord{X: 1, Y: 2},
			Body:   []Coord{{X: 1, Y: 2}, {X: 1, Y: 3}}},
		{
			ID:     "2",
			Health: 100,
			Head:   Coord{X: 3, Y: 2},
			Body:   []Coord{{X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}}}}

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

type idk []struct {
	X int
	Y int
}

func getGameState1Result() coreplayer.GameBoard {
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
		Height:  4,
		Width:   4,
		Food:    food,
		Hazards: hazards,
		Snakes:  snakes,
	}
}

func jsonCompare[T GameState | coreplayer.GameBoard](a T, b T) bool {
	ajson, _ := json.Marshal(a)
	bjson, _ := json.Marshal(b)
	return bytes.Compare(ajson, bjson) == 0
}

func printJsonStruct[T GameState | coreplayer.GameBoard](a T) {
	data, _ := json.Marshal(a)
	println(string(data))
}

func Test_stateConverter(t *testing.T) {
	conv := minimaxplayer.StateConverterImpl{}
	var res GameState
	res = conv.CoreToState(getGameState1Result())
	var expected GameState
	expected = getGameStateTest1()
	if !jsonCompare(res, expected) {
		println("Could not convert test 1 to test2, equality failed")
		println("res: ")
		printJsonStruct(res)
		println("expected: ")
		printJsonStruct(expected)
		t.FailNow()
	}

}
