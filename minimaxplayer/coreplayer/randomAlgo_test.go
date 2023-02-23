package coreplayer_test

import (
	"jacksnake/minimaxplayer/coreplayer"
	"testing"
)

type FakeSimulator struct {
	Moves []string
}

func (*FakeSimulator) SimulateMoves(board coreplayer.GameBoard, moves []coreplayer.SnakeMove) coreplayer.GameBoard {
	return board
}

func (sim *FakeSimulator) GetValidMoves(board coreplayer.GameBoard, snakeId string) []string {
	return sim.Moves
}

func Test_random(t *testing.T) {
	random := coreplayer.RandomAlgo{}
	sim := FakeSimulator{}
	sim.Moves = []string{"up", "down"}
	random.Init(&sim)
	move := random.Move(coreplayer.GameBoard{}, "1")
	if move != "up" && move != "down" {
		t.Fail()
	}
}

func Test_random1(t *testing.T) {
	random := coreplayer.RandomAlgo{}
	sim := FakeSimulator{}
	sim.Moves = []string{"up"}
	random.Init(&sim)
	move := random.Move(coreplayer.GameBoard{}, "1")
	if move != "up" {
		t.Fail()
	}
}
