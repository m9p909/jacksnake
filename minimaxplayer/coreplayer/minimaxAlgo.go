package coreplayer

import (
	"encoding/json"
	"math"
)

type MinimaxAlgoMove struct {
	simulator   Simulator
	evaluator   Evaluator
	maxDepth    int
	playerIndex int
	playerId    string
}

func NewMinimaxAlgoMove(sim Simulator, eval Evaluator, maxDepth int) *MinimaxAlgoMove {
	return &MinimaxAlgoMove{
		simulator: sim,
		evaluator: eval,
		maxDepth:  maxDepth,
	}
}

func findSnakeById(snakes *[]Snake, id string) (int, *Snake) {
	for i, snake := range *snakes {
		if snake.ID == id {
			return i, &snake
		}
	}
	println("NO SNAKE FOUND")
	return -1, nil
}

func movePlayerSnakeToFront(snakes *[]Snake, id string) *[]Snake {
	for index, snake := range *snakes {
		if snake.ID == id {
			first := (*snakes)[0]
			(*snakes)[0] = (*snakes)[index]
			(*snakes)[index] = first
			return snakes
		}
	}
	return nil
}

func (minimax *MinimaxAlgoMove) Move(board GameBoard, snakeId string) string {
	println(snakeId)
	data, _ := json.Marshal(board)
	println(string(data))
	var snek *Snake
	minimax.playerIndex, snek = findSnakeById(&board.Snakes, snakeId)
	minimax.playerId = snek.ID
	res := minimax.startMinimax(&board)
	return res
}

func makeNewSnakeMoves(board *GameBoard) []SnakeMove {
	return make([]SnakeMove, len(board.Snakes))
}

func (minimax *MinimaxAlgoMove) startMinimax(board *GameBoard) string {
	moves := minimax.simulator.GetValidMoves(*board, minimax.playerId)
	max := 0.0
	bestMove := "down"
	for _, move := range moves {
		nextMoves := makeNewSnakeMoves(board)
		nextMoves[minimax.playerIndex] = SnakeMove{ID: minimax.playerId, Move: move}
		score := minimax.runMinimax(board, getNextSnakeIndex(board, minimax.playerIndex), 1, nextMoves)
		if score > max {
			max = score
			bestMove = move
		}
	}
	return bestMove
}

func getNextSnakeIndex(board *GameBoard, currentSnake int) int {
	return (currentSnake + 1) % len(board.Snakes)
}

func (minimax *MinimaxAlgoMove) runMinimax(board *GameBoard, snakeIndex int, depth int, moves []SnakeMove) float64 {
	if depth > minimax.maxDepth {
		return minimax.evaluator.EvaluateBoard(board, minimax.playerId)
	}

	//println(snakeIndex)
	if snakeIndex == minimax.playerIndex {
		// print the moves structure
		//fmt.Printf("moves: %v\n", moves)
		newBoard := minimax.simulator.SimulateMoves(*board, moves)
		max := 0.0
		validMoves := minimax.simulator.GetValidMoves(newBoard, minimax.playerId)
		for _, move := range validMoves {
			moves := makeNewSnakeMoves(board)
			moves[minimax.playerIndex] = SnakeMove{ID: minimax.playerId, Move: move}
			score := minimax.runMinimax(&newBoard, getNextSnakeIndex(board, snakeIndex), depth+1, moves)
			if score > max {
				max = score
			}
		}
		return max
	} else {
		snakeId := board.Snakes[snakeIndex].ID
		min := math.Inf(1)
		validMoves := minimax.simulator.GetValidMoves(*board, snakeId)
		for _, move := range validMoves {
			// append works correctly
			newMoves := makeNewSnakeMoves(board)
			copy(newMoves, moves)
			newMoves[snakeIndex] = SnakeMove{ID: snakeId, Move: move}
			score := minimax.runMinimax(board, getNextSnakeIndex(board, snakeIndex), depth, newMoves)
			if score < min {
				min = score
			}
		}
		return min
	}
}
