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
	/*	println(snakeId)
		data, _ := json.Marshal(board)
		println(string(data))*/
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
	// d, _ := json.Marshal(board)
	// println(string(d))
	ch := make(chan MoveResult)

	for i, move := range moves {
		go func(i int, move Direction, board GameBoard) {
			nextMoves := makeNewSnakeMoves(&board)
			nextMoves[ids.playerIndex] = SnakeMove{ID: ids.playerId, Move: move}
			score := minimax.runMinimax(minimaxArgs{&board, getNextSnakeIndex(&board, ids.playerIndex), 1, nextMoves, ids, math.Inf(-1), math.Inf(1)})
			result := MoveResult{move: move, score: score}
			ch <- result
		}(i, move, *board)
	}

	for i := range scores {
		score := <-ch
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
	// println(DirectionToString(bestMove))

	return bestMove
}

func getNextSnakeIndex(board *GameBoard, currentSnake uint8) uint8 {
	i := (uint8(currentSnake) + 1) % uint8(len(board.Snakes))
	return i
}

type minimaxArgs struct {
	board      *GameBoard
	snakeIndex uint8
	depth      int
	moves      []SnakeMove
	ids        *PlayerIds
	alpha      float64
	beta       float64
}

func (minimax *MinimaxAlgoMove) getBestMove(args *minimaxArgs) float64 {
	max := 0.0
	validMoves := minimax.simulator.GetValidMoves(args.board, args.ids.playerId)
	for _, move := range validMoves {
		moves := makeNewSnakeMoves(args.board)
		moves[args.ids.playerIndex] = SnakeMove{ID: args.ids.playerId, Move: move}
		score := minimax.runMinimax(minimaxArgs{args.board, getNextSnakeIndex(args.board, args.snakeIndex), args.depth + 1, moves, args.ids, args.alpha, args.beta})
		if score > args.beta {
			break
		}
		if score > max {
			max = score
		}
		if score > args.alpha {
			args.alpha = score
		}
	}
	return max
}

func (minimax *MinimaxAlgoMove) doMinimizingPlayer(args *minimaxArgs) float64 {
	snakeId := args.board.Snakes[args.snakeIndex].ID
	min := math.Inf(1)
	validMoves := minimax.simulator.GetValidMoves(args.board, snakeId)
	for _, move := range validMoves {
		newMoves := makeNewSnakeMoves(args.board)
		copy(newMoves, args.moves)
		newMoves[args.snakeIndex] = SnakeMove{ID: snakeId, Move: move}
		score := minimax.runMinimax(minimaxArgs{args.board, getNextSnakeIndex(args.board, args.snakeIndex), args.depth, newMoves, args.ids, args.alpha, args.beta})
		if score < args.alpha {
			break
		}
		if score < min {
			min = score
		}
		if score < args.beta {
			args.beta = score
		}
	}
	return min
}

func (minimax *MinimaxAlgoMove) runMinimax(args minimaxArgs) float64 {
	if args.depth > minimax.maxDepth {
		eval := minimax.evaluator.EvaluateBoard(args.board, args.ids.playerId)
		return eval
	}
	// println(snakeIndex)
	if args.snakeIndex == args.ids.playerIndex {
		// print the moves structure
		// fmt.Printf("moves: %v\n", moves)
		newBoard := args.board.Clone()
		minimax.simulator.SimulateMoves(&newBoard, args.moves)
		// prune the branch because health is 0
		if newBoard.Snakes[args.ids.playerIndex].Health == 0 {
			// println("found dead path")
			return minimax.evaluator.EvaluateBoard(args.board, args.ids.playerId) * 0.1 // should be super weak
		}
		args.board = &newBoard
		res := minimax.getBestMove(&args)
		return res
	} else {
		return minimax.doMinimizingPlayer(&args)
	}
}
