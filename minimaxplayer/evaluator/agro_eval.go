package evaluator

import (
	"fmt"
	. "jacksnake/minimaxplayer/coreplayer"
)

type AgroEvaluator struct{}

func NewAgroEvaluator() Evaluator {
	return &AgroEvaluator{}
}

func lengthScore(board *GameBoard, snakeId SnakeID) float64 {
	max := 0

	for _, snake := range board.Snakes {
		if(snake.Health > 0 && len(snake.Body) > max)  {
			max = len(snake.Body)
		}
	}
	return float64(len(board.Snakes[snakeId].Body))/float64(max)
}





func (*AgroEvaluator) EvaluateBoard(board *GameBoard, snakeId SnakeID, complete bool, count int) float64 {
	snake := findSnakeById(&board.Snakes, snakeId)
	if snake != nil {
		healthScore := getHealthScore(snake)
		if healthScore < 0 || healthScore > 1 {
			fmt.Println(healthScore)
			panic("bad health score")
		}
		otherSnakesHealth := getOtherSnakeHealthScore(board, snake)
		spaceScore := evaluateSpaceConstraint(board, snakeId)
		score := otherSnakesHealth * 0.2+ healthScore*0.4 + spaceScore * 0.2 + lengthScore(board, snakeId) * 0.2
		// if the max depth is reached
		if score < 0 {
			println("neg score")
		}

		if score > 1 {
			println("score too big")
		}
		// reduce weight of score if not at end of game
		if !complete {
			score = score * 0.01 * (float64(count) + 1)
		}
		if score <= 0 || score > 1 {
			panic("Invalid score")
		}
		return score
	}
	// println("no snake found, this should never happen")
	return 0
}
