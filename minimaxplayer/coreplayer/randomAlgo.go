package coreplayer

import (
	"math/rand"
)

type RandomAlgo struct {
	simulator Simulator
}

func (rand *RandomAlgo) Init(simulator Simulator) {
	rand.simulator = simulator
}

func (minimax *RandomAlgo) Move(board GameBoard, youId string) string {
	safeMoves := minimax.simulator.GetValidMoves(board, youId)
	if len(safeMoves) <= 0 {
		println("NO safe moves detected")
		return "down"
	}
	return safeMoves[rand.Intn(len(safeMoves))]
}
