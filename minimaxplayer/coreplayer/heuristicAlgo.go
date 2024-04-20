package coreplayer

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type HeuristicAlgo struct {
	sim  Simulator
	eval Evaluator
}

func NewHeuristicAlgo(sim Simulator, eval Evaluator) *HeuristicAlgo {
	return &HeuristicAlgo{sim, eval}
}

func (h *HeuristicAlgo) Move(board GameBoard, snakeId SnakeID) Direction {
	return h.scoreAllPossibleGameBoards(&board, snakeId)
}

func (h *HeuristicAlgo) scoreAllPossibleGameBoards(board *GameBoard, snakeId SnakeID) Direction {
	d, _ := json.Marshal(board)
	println(string(d))

	left := 0.0
	right := 00.0
	up := 00.0
	down := 00.0
	for _, combination := range BuildProductOfDirections(len(board.Snakes)) {
		// shuffle
		for i := range combination {
			j := rand.Intn(i + 1)
			combination[i], combination[j] = combination[j], combination[i]
		}
		boardCopy := board.Clone()
		moves := func() []SnakeMove {
			var res []SnakeMove
			for _, snake := range board.Snakes {
				res = append(res, SnakeMove{snake.ID, combination[snake.ID]})
			}
			return res
		}()
		h.sim.SimulateMoves(&boardCopy, moves)
		if boardCopy.Snakes[snakeId].Health == 0 {
			continue
		}
		score := h.eval.EvaluateBoard(&boardCopy, snakeId, true, 4)
		dir := combination[snakeId]
		if dir == LEFT {
			left = left + score
		}
		if dir == RIGHT {
			right = right + score
		}
		if dir == UP {
			up = up + score
		}
		if dir == DOWN {
			down = down + score
		}
	}
	fmt.Printf("\n Scores: %f %f %f %f \n", left, right, up, down)
	max := left
	dir := LEFT
	if right > max {
		max = right
		dir = RIGHT
	}
	if up > max {
		max = up
		dir = UP
	}
	if down > max {
		max = down
		dir = DOWN
	}
	return dir
}

func BuildProductOfDirections(i int) [][]Direction {
	combinations := [][]Direction{}
	var helper func([]Direction)

	helper = func(dirs []Direction) {
		temp := make([]Direction, len(dirs))
		copy(temp, dirs)
		if len(temp) == i {
			combinations = append(combinations, temp)
			return
		}
		for _, dir := range []Direction{LEFT, RIGHT, UP, DOWN} {
			helper(append(temp, dir))
		}
	}

	helper([]Direction{})
	return combinations

}
