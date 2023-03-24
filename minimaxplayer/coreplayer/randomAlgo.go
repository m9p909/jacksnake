package coreplayer

import (
	"math/rand"
)

type RandomAlgo struct {
	simulator Simulator
}

func NewRandomAlgo(simulator Simulator) Player {
	return &RandomAlgo{
		simulator: simulator,
	}
}

func (minimax *RandomAlgo) Move(board GameBoard, youId SnakeID) Direction {
	safeMoves := minimax.simulator.GetValidMoves(&board, youId)
	if len(safeMoves) <= 0 {
		println("NO safe moves detected")
		return DOWN
	}
	return safeMoves[rand.Intn(len(safeMoves))]
}

func (algo *RandomAlgo) Clone() RandomAlgo {
	return *algo
}
