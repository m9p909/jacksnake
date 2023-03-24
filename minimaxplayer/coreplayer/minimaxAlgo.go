package coreplayer

import (
	"math"
)

type PlayerIds struct {
	playerIndex uint8
	playerId    SnakeID
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

func findSnakeById(snakes []Snake, id SnakeID) (uint8, *Snake) {
	for i := range snakes {
		if snakes[i].ID == id {
			return uint8(i), &snakes[i]
		}
	}
	println("snake not found")
	return 0, nil
}

func (minimax *MinimaxAlgoMove) Move(board GameBoard, snakeId SnakeID) Direction {
	// for collecting test data
	/*
		println(snakeId)
		data, _ := json.Marshal(board)
		println(string(data))
	*/
	playerIds := PlayerIds{}
	var snek *Snake
	playerIds.playerIndex, snek = findSnakeById(board.Snakes, snakeId)
	playerIds.playerId = snek.ID
	res := minimax.startMinimax(&board, &playerIds)
	return res
}

func makeNewSnakeMoves(board *GameBoard) []SnakeMove {
	return make([]SnakeMove, len(board.Snakes))
}

type MoveResult struct {
	score float64
	move  Direction
}

func (minimax *MinimaxAlgoMove) getScores(moves []Direction, board *GameBoard, ids *PlayerIds) []MoveResult {
	scores := make([]MoveResult, len(moves))
	chans := make([]chan MoveResult, len(moves))
	for i := range chans {
		chans[i] = make(chan MoveResult)
	}
	for i, move := range moves {
		nextMoves := makeNewSnakeMoves(board)
		nextMoves[ids.playerIndex] = SnakeMove{ID: ids.playerId, Move: move}
		go func(index int, move Direction) {
			newBoard := *board
			score := minimax.runMinimax(&newBoard, getNextSnakeIndex(&newBoard, ids.playerIndex), 1, nextMoves, ids)
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

func (minimax *MinimaxAlgoMove) startMinimax(board *GameBoard, ids *PlayerIds) Direction {
	moves := minimax.simulator.GetValidMoves(board, ids.playerId)
	max := 0.0
	bestMove := DOWN
	scores := minimax.getScores(moves, board, ids)
	for _, score := range scores {
		if score.score > max {
			max = score.score
			bestMove = score.move
		}
	}
	println(DirectionToString(bestMove))

	return bestMove
}

func getNextSnakeIndex(board *GameBoard, currentSnake uint8) uint8 {
	i := (uint8(currentSnake) + 1) % uint8(len(board.Snakes))
	return i
}

func (minimax *MinimaxAlgoMove) runMinimax(board *GameBoard, snakeIndex uint8, depth int, moves []SnakeMove, ids *PlayerIds) float64 {
	if depth > minimax.maxDepth {
		eval := minimax.evaluator.EvaluateBoard(board, ids.playerId)
		return eval
	}

	// println(snakeIndex)
	if snakeIndex == ids.playerIndex {
		// print the moves structure
		// fmt.Printf("moves: %v\n", moves)
		newBoard := *board
		minimax.simulator.SimulateMoves(&newBoard, moves)
		if newBoard.Snakes[snakeIndex].Health == 0 {
			return 0
		}
		max := 0.0
		validMoves := minimax.simulator.GetValidMoves(&newBoard, ids.playerId)
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
		validMoves := minimax.simulator.GetValidMoves(board, snakeId)
		for _, move := range validMoves {
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
