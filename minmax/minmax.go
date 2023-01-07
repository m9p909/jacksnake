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

type possibleGameState struct {
	state GameState
	move  string
}

func getPossibleStates(state GameState, turn int) []possibleGameState {
	//TODO
	return []possibleGameState{}

}

func (minmax *MinMaxImpl) evaluateState(state GameState) float64 {
	// TODO
	return 0
}

type MinimaxDataInput struct {
	depth    int
	playerId int
	alpha    float64
	beta     float64
	state    GameState
}

type MinimaxDataOutput struct {
	value float64
	turn  string
}

func (minmax *MinMaxImpl) minimax(input MinimaxDataInput) MinimaxDataOutput {

	if input.depth == minmax.maxDepth {
		state := minmax.evaluateState(input.state)
		return MinimaxDataOutput{
			value: state,
			turn:  "unknown",
		}
	}

	if input.playerId == 0 {
		bestVal := math.Inf(-1)
		bestState := possibleGameState{}
		states := getPossibleStates(input.state, input.playerId)
		for _, possibleState := range states {

			minimaxOutput := minmax.minimax(MinimaxDataInput{
				input.depth + 1,
				input.playerId + 1,
				input.alpha,
				input.beta,
				possibleState.state})

			if minimaxOutput.value > bestVal {
				bestVal = minimaxOutput.value
				bestState = possibleState
			}

			input.alpha = max(input.alpha, bestVal)
			if input.beta <= input.alpha {
				break
			}
		}
		return MinimaxDataOutput{
			value: bestVal,
			turn:  bestState.move,
		}

	} else {
		worstVal := math.Inf(1)
		worstState := possibleGameState{}
		states := getPossibleStates(input.state, input.playerId)
		for _, possibleState := range states {

			minimaxOutput := minmax.minimax(MinimaxDataInput{
				input.depth + 1,
				input.playerId + 1,
				input.alpha,
				input.beta,
				possibleState.state})

			if minimaxOutput.value < worstVal {
				worstVal = minimaxOutput.value
				worstState = possibleState
			}

			input.beta = min(input.beta, worstVal)
			if input.beta <= input.alpha {
				break
			}
		}

		return MinimaxDataOutput{
			value: worstVal,
			turn:  worstState.move,
		}
	}

}
