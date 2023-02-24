package coreplayer

import (
	"math"
)

type PlayerIds struct {
	playerIndex int
	playerId    string
}

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
	// for collecting test data
	/*
		println(snakeId)
		data, _ := json.Marshal(board)
		println(string(data))
	*/
	playerIds := PlayerIds{}
	var snek *Snake
	playerIds.playerIndex, snek = findSnakeById(&board.Snakes, snakeId)
	playerIds.playerId = snek.ID
	res := minimax.startMinimax(&board, &playerIds)
	return res
}

func makeNewSnakeMoves(board *GameBoard) []SnakeMove {
	return make([]SnakeMove, len(board.Snakes))
}

type MoveResult struct {
	score float64
	move  string
}

func (minimax *MinimaxAlgoMove) getScores(moves []string, board *GameBoard, ids *PlayerIds) []MoveResult {
	scores := make([]MoveResult, len(moves))
	chans := make([]chan MoveResult, len(moves))
	for i := range chans {
		chans[i] = make(chan MoveResult)
	}
	for i, move := range moves {
		nextMoves := makeNewSnakeMoves(board)
		nextMoves[ids.playerIndex] = SnakeMove{ID: ids.playerId, Move: move}
		go func(index int, move string) {
			score := minimax.runMinimax(board, getNextSnakeIndex(board, ids.playerIndex), 1, nextMoves, ids)
			result := MoveResult{move: move, score: score}
			chans[index] <- result
		}(i, move)
	}
	for i := range chans {
		score := <-chans[i]
		scores[i] = score
	}
	return scores
}

func (minimax *MinimaxAlgoMove) startMinimax(board *GameBoard, ids *PlayerIds) string {
	moves := minimax.simulator.GetValidMoves(*board, ids.playerId)
	max := 0.0
	bestMove := "down"
	scores := minimax.getScores(moves, board, ids)
	for _, score := range scores {
		if score.score > max {
			max = score.score
			bestMove = score.move
		}
	}

	return bestMove
}

func getNextSnakeIndex(board *GameBoard, currentSnake int) int {
	return (currentSnake + 1) % len(board.Snakes)
}

func (minimax *MinimaxAlgoMove) runMinimax(board *GameBoard, snakeIndex int, depth int, moves []SnakeMove, ids *PlayerIds) float64 {
	if depth > minimax.maxDepth {
		return minimax.evaluator.EvaluateBoard(board, ids.playerId)
	}

	// println(snakeIndex)
	if snakeIndex == ids.playerIndex {
		// print the moves structure
		// fmt.Printf("moves: %v\n", moves)
		newBoard := minimax.simulator.SimulateMoves(*board, moves)
		max := 0.0
		validMoves := minimax.simulator.GetValidMoves(newBoard, ids.playerId)
		for _, move := range validMoves {
			moves := makeNewSnakeMoves(board)
			moves[ids.playerIndex] = SnakeMove{ID: ids.playerId, Move: move}
			score := minimax.runMinimax(&newBoard, getNextSnakeIndex(board, snakeIndex), depth+1, moves, ids)
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
			score := minimax.runMinimax(board, getNextSnakeIndex(board, snakeIndex), depth, newMoves, ids)
			if score < min {
				min = score
			}
		}
		return min
	}
}
