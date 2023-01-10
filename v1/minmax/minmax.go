package minmax

import (
	"errors"
	"jacksnake/evaluateBoard"
	. "jacksnake/models"
	"jacksnake/safemoves"
	"jacksnake/simulation"
	"math"
)

// Snake IDS

func getSnakes(state GameState) []Battlesnake {
	return state.Board.Snakes
}

func getNumSnakes(state GameState) int {
	return len(state.Board.Snakes)
}

type MiniMax struct {
	maxDepth     int
	numSnakes    int
	initialState GameState
	youID        int
}

func findSnakeWithId(snakes []Battlesnake, id string) (int, error) {
	for index, snake := range snakes {
		if snake.ID == id {
			return index, nil
		}
	}
	return -1, errors.New("could not find snake")
}

func NewMiniMax(maxDepth int, initialState GameState) MiniMax {
	id, err := findSnakeWithId(
		initialState.Board.Snakes,
		initialState.You.ID)
	if err != nil {
		println("Is the player in this game?")
	}
	return MiniMax{
		maxDepth,
		getNumSnakes(initialState),
		initialState,
		id,
	}
}

func (minmax *MiniMax) Analyze() string {
	decision := minmax.minimax(MinimaxDataInput{
		minmax.youID,
		math.Inf(1),
		math.Inf(-1),
		minmax.initialState,
		0,
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
		state := simulation.SimulateMoveBySnake(state, move, snakes[turn])
		output = append(output, possibleGameState{
			state: state,
			move:  move,
		})
	}

	if len(moves) == 0 {
		state.Board.Snakes[turn] = Battlesnake{
			ID:     state.Board.Snakes[turn].ID,
			Health: 0,
			Body:   []Coord{},
			Head:   Coord{},
		}

		output = []possibleGameState{
			{
				state,
				"dead",
			},
		}
	}
	return output

}

func (minmax *MiniMax) evaluateState(state GameState) float64 {
	value, err := evaluateboard.EvaluatePossibleState(state)
	if err != nil {
		return 0.5 // return neutral value.

	}
	return value
}

type MinimaxDataInput struct {
	playerId int
	alpha    float64
	beta     float64
	state    GameState
	depth    int
}

type MinimaxDataOutput struct {
	value float64
	turn  string
}

func (minimax *MiniMax) minimax(input MinimaxDataInput) MinimaxDataOutput {

	if input.depth == minimax.maxDepth {
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
			if possibleState.move == "dead" {
				break
			}

			minimaxOutput := minimax.minimax(MinimaxDataInput{
				(input.playerId + 1) % minimax.numSnakes,
				input.alpha,
				input.beta,
				possibleState.state,
				input.depth + 1,
			})

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
				(input.playerId + 1) % minimax.numSnakes,
				input.alpha,
				input.beta,
				possibleState.state,
				input.depth + 1})

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
