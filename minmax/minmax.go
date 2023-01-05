package minmax

import (
	. "jacksnake/models"
	"math"
)

type MinMax interface {
	Analyze(depth int, heuristicAlgorithm func(state GameState) float64) string
}

type MinMaxImpl struct {
	maxDepth int
}

func (minmax *MinMaxImpl) Analyze(depth int) {

}

type Node struct {
	children []*Node
}

func max(a float64, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a float64, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func (minmax *MinMaxImpl) evaluateState(state GameState) float64 {
	return 0
}

func getPossibleStates(state GameState) []*GameState {

}

func (minmax *MinMaxImpl) minimax(depth int, isMaximizingPlayer bool, alpha float64, beta float64, state GameState) float64 {

	if depth == minmax.maxDepth {
		return minmax.evaluateState(state)
	}

	if isMaximizingPlayer {
		bestVal := math.Inf(-1)
		for _, child := range node.children {
			value := minimax(child, depth+1, false, alpha, beta)
			bestVal = max(bestVal, value)
			alpha = max(alpha, bestVal)
			if beta <= alpha {
				break
			}
		}
		return bestVal
	} else {
		bestVal := math.Inf(1)
		for _, child := range node.children {
			value := minimax(child, depth+1, true, alpha, beta)
			bestVal = min(bestVal, value)
			beta = min(beta, bestVal)
			if beta <= alpha {
				break
			}
		}
		return bestVal
	}

}
