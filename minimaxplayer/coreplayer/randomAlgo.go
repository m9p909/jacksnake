package coreplayer

import (
	"math/rand"
)

type RandomAlgo struct {
	simulator Simulator
}

func (minimax *RandomAlgo) Move(board GameBoard, youId string) string {
	safeMoves := minimax.simulator.GetValidMoves(board, youId)
	return safeMoves[rand.Intn(len(safeMoves))]
}
