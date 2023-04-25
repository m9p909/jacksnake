package customsimulator_test

import (
	"encoding/json"
	. "jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/customsimulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestBoard() GameBoard {
	/*
		10  F            0 1
		9                0 1
		8                0 1
		7
		6
		5
		4
		3
		2
		1
		0
		 0 1 2 3 4 5 6 7 8 9 10
	*/

	boardjson := "{\"Turn\":0,\"Height\":11,\"Width\":11,\"Food\":[{\"X\":8,\"Y\":10},{\"X\":8,\"Y\":9},{\"X\":8,\"Y\":8}],\"Snakes\":[{\"ID\":0,\"Body\":[{\"X\":8,\"Y\":10},{\"X\":8,\"Y\":9},{\"X\":8,\"Y\":8}],\"Health\":99},{\"ID\":1,\"Body\":[{\"X\":9,\"Y\":10},{\"X\":9,\"Y\":9},{\"X\":9,\"Y\":8}],\"Health\":99}],\"Hazards\":[]}"
	var res GameBoard
	json.Unmarshal([]byte(boardjson), &res)
	return res
}

func TestSimulationMutatesBoard(t *testing.T) {
	res := getTestBoard()
	res2 := res.Clone()
	sim := customsimulator.New()
	moves1 := sim.GetValidMoves(&res, 0)
	moves2 := sim.GetValidMoves(&res, 1)
	moves3 := []SnakeMove{
		{
			ID:   0,
			Move: moves1[0],
		},
		{ID: 1, Move: moves2[0]},
	}

	sim.SimulateMoves(&res, moves3)
	assert.NotEqualValues(t, res, res2)
	assert.NotEqualValues(t, res.Snakes[0].Body, res2.Snakes[0].Body)
}

func TestActuallyWorks(t *testing.T) {
	/*
		10  F            0 1 1
		9                0 1
		8
		7
		6
		5
		4
		3
		2
		1
		0
		 0 1 2 3 4 5 6 7 8 9 10
	*/
	board := getTestBoard()
	moves := []SnakeMove{
		{0, RIGHT},
		{1, RIGHT},
	}

	customsimulator.New().SimulateMoves(&board, moves)
	assert.EqualValues(t, []Point{{10, 10}, {9, 10}, {9, 9}}, board.Snakes[1].Body)
	assert.EqualValues(t, []Point{{9, 10}, {8, 10}, {8, 9}}, board.Snakes[0].Body)
	assert.Equal(t, board.Snakes[0].Health, uint8(0))
}
