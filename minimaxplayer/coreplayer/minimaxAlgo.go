package coreplayer

import (
	"math"
)

type MinimaxAlgoMove struct {
	simulator Simulator
	evaluator Evaluator
	maxDepth  int
}

func NewMinimaxAlgoMove(sim Simulator, eval Evaluator, maxDepth int) *MinimaxAlgoMove {
	return &MinimaxAlgoMove{
		simulator: sim,
		evaluator: eval,
		maxDepth:  maxDepth,
	}
}

func findSnakeById(snakes *[]Snake, id string) *Snake {
	for _, snake := range *snakes {
		if snake.ID == id {
			return &snake
		}
	}
	return nil
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
	movePlayerSnakeToFront(&board.Snakes, snakeId)
	res := minimax.startMinimax(&board)
	return res
}

func (minimax *MinimaxAlgoMove) startMinimax(board *GameBoard) string {
	playerID := board.Snakes[0].ID
	moves := minimax.simulator.GetValidMoves(*board, playerID)
	max := 0.0
	bestMove := "down"
	for _, move := range moves {
		score := minimax.runMinimax(board, getNextSnake(board, 0), 1, &[]SnakeMove{{ID: playerID, Move: move}})
		if score > max {
			max = score
			bestMove = move
		}
	}
	return bestMove
}

func getNextSnake(board *GameBoard, currentSnake int) int {
	return (currentSnake + 1) % len(board.Snakes)
}

func (minimax *MinimaxAlgoMove) runMinimax(board *GameBoard, snake int, depth int, moves *[]SnakeMove) float64 {
	if depth > minimax.maxDepth {
		return minimax.evaluator.EvaluateBoard(board, board.Snakes[0].ID)
	}
	if snake == 0 {
		mainSnakeId := board.Snakes[0].ID
		newBoard := minimax.simulator.SimulateMoves(*board, *moves)
		max := 0.0
		validMoves := minimax.simulator.GetValidMoves(newBoard, mainSnakeId)
		for _, move := range validMoves {
			score := minimax.runMinimax(&newBoard, getNextSnake(board, snake), depth+1, &[]SnakeMove{{ID: mainSnakeId, Move: move}})
			if score > max {
				max = score
			}
		}
		return max
	} else {
		snakeId := board.Snakes[snake].ID
		min := math.Inf(1)
		validMoves := minimax.simulator.GetValidMoves(*board, snakeId)
		nextSnake := getNextSnake(board, snake)
		for _, move := range validMoves {
			currentmoves := append(*moves, SnakeMove{ID: snakeId, Move: move}) // possible bug, not sure if append modifies moves
			score := minimax.runMinimax(board, nextSnake, depth, &currentmoves)
			if score < min {
				min = score
			}
		}
		return min
	}
}
