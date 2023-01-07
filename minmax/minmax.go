package minmax

import (
	"jacksnake/evaluateBoard"
	. "jacksnake/models"
	"jacksnake/safemoves"
	"jacksnake/simulation"
	"math"
)

// Snake IDS

func getSnakes(state GameState) []Battlesnake {
	numSnakes := len(state.Board.Snakes) + 1
	snakes := make([]Battlesnake, numSnakes)
	for index, snek := range state.Board.Snakes {
		snakes[index] = snek
	}
	snakes[numSnakes-1] = state.You
	return snakes
}

func getNumSnakes(state GameState) int {
	return len(state.Board.Snakes) + 1
}

type MiniMax struct {
	maxDepth     int
	numSnakes    int
	initialState GameState
	youID        int
}

func NewMiniMax(maxDepth int, initialState GameState) MiniMax {
	return MiniMax{
		maxDepth,
		getNumSnakes(initialState),
		initialState,
		getNumSnakes(initialState) - 1,
	}
}

func (minmax *MiniMax) Analyze() string {
	decision := minmax.minimax(MinimaxDataInput{
		minmax.numSnakes - 1,
		math.Inf(1),
		math.Inf(-1),
		minmax.initialState,
	})
	return decision.turn
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

// int i = 0 is the player. others are other players
func getPossibleStates(state GameState, turn int) []possibleGameState {
	output := []possibleGameState{}
	snakes := getSnakes(state)

	moves := safemoves.GetSafeMovesBySnake(state, snakes[turn])
	for _, move := range moves {
		state := simulation.SimulateMove(state, move)
		output = append(output, possibleGameState{
			state: state,
			move:  move,
		})
	}
	return output

}

func (minmax *MiniMax) evaluateState(state GameState) float64 {
	return evaluateboard.EvaluatePossibleState(state)
}

type MinimaxDataInput struct {
	playerId int
	alpha    float64
	beta     float64
	state    GameState
}

type MinimaxDataOutput struct {
	value float64
	turn  string
}

func (minimax *MiniMax) minimax(input MinimaxDataInput) MinimaxDataOutput {

	if minimax.maxDepth == minimax.maxDepth {
		state := minimax.evaluateState(input.state)
		return MinimaxDataOutput{
			value: state,
			turn:  "unknown",
		}
	}

	if input.playerId == minimax.youID {
		bestVal := math.Inf(-1)
		bestState := possibleGameState{}
		states := getPossibleStates(input.state, input.playerId)
		for _, possibleState := range states {

			minimaxOutput := minimax.minimax(MinimaxDataInput{
				0,
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

			minimaxOutput := minimax.minimax(MinimaxDataInput{
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
