package coreplayer

import (
	"errors"
)

type MinimaxAlgoMove struct {
	simulator Simulator
	evaluator Evaluator
	maxDepth  int64
}

func getMinimaxAlgoMove(sim Simulator, eval Evaluator) *MinimaxAlgoMove {
	return &MinimaxAlgoMove{
		simulator: sim,
		evaluator: eval,
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
}

func (minimax *MinimaxAlgoMove) Move(board *GameBoard, snakeId string) (error, string) {
	movePlayerSnakeToFront(&board.Snakes, snakeId)
	if playerSnake == nil {
		return errors.New("could not find snake"), "down"
	}
	return nil, ""
}

func (minimax *MinimaxAlgoMove) startMinimax(board *GameBoard, snake int) string {
	playerID := board.Snakes[0].ID
	moves := minimax.simulator.GetValidMoves(*board, playerID)
	max := 0
	move := "down"
	for _, move := range moves {
		score := minimax.runMinimax(board, 1, 1, 0, []SnakeMove{{ID: playerID, Move: move}})
		if score > max {
			max = score
			move = move
		}
	}
	return move
}

func (minimax *MinimaxAlgoMove) runMinimax(board *GameBoard, snake int, depth int, score float64, moves []SnakeMove) int {
	if depth > int(minimax.maxDepth) {
		return score
	}
	if snake == 0 {
	}
}
