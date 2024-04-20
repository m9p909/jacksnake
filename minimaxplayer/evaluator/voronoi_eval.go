package evaluator

import (
	"github.com/emirpasic/gods/queues/circularbuffer"
	. "jacksnake/minimaxplayer/coreplayer"
)

type VoronoiEval struct{}

func NewVoronoiEval() Evaluator {
	return &VoronoiEval{}
}

type voronoiIndex struct {
	closestSnake SnakeID
	dist         uint8
	snakeId      SnakeID
}

func makeEmptyVoronoiBoard(height uint8, width uint8) [][]*voronoiIndex {
	board := make([][]*voronoiIndex, height)
	for b := range board {
		row := make([]*voronoiIndex, width)
		for val := range row {
			row[val] = &voronoiIndex{closestSnake: 255,
				dist:    255,
				snakeId: 255,
			}
		}
		board[b] = row
	}

	return board
}

func buildVoronoiSnakeBoard(snakes []Snake, height uint8, width uint8) [][]*voronoiIndex {
	board := makeEmptyVoronoiBoard(height, width)

	for _, snake := range snakes {
		if snake.Health > 0 {
			for _, body := range snake.Body {
				board[body.Y][body.X] = &voronoiIndex{
					closestSnake: 255,
					dist:         255,
					snakeId:      snake.ID,
				}
			}
		}
	}

	return board
}

type qNode struct {
	p    Point
	dist uint8
}

func VoronoiScore(board *GameBoard) []float64 {
	arrBoard := buildVoronoiSnakeBoard(board.Snakes, board.Height, board.Width)
	for _, snake := range board.Snakes {
		if snake.Health <= 0 {
			continue
		}
		visited := make(map[Point]bool, 20)
		hasVisted := func(p Point) bool {
			_, ok := visited[p]
			return ok
		}
		head := snake.Body[0]
		qPoint := circularbuffer.New(int(board.Height+board.Width) * 2)
		enqueueNextPoints(board, qNode{head, 0}, hasVisted, qPoint, visited)
		for qPoint.Size() > 0 {
			p, ok := qPoint.Dequeue()
			if !ok {
				panic("queue should never be empty")
			}
			pvalue := p.(qNode)

			current := arrBoard[pvalue.p.Y][pvalue.p.X]
			if !(current.snakeId == 255) {
				continue
			}

			if pvalue.dist < current.dist {
				current.dist = pvalue.dist
				current.closestSnake = snake.ID
			} else if pvalue.dist == current.dist {
				current.dist = pvalue.dist
				// nobody gets it
				current.closestSnake = 67
			}

			// determine nexxt point

			enqueueNextPoints(board, pvalue, hasVisted, qPoint, visited)
		}
	}
	var results = make([]float64, 4)
	for _, snake := range board.Snakes {

		sum := 0

		for _, row := range arrBoard {
			for _, cell := range row {
				if cell.closestSnake == snake.ID {
					sum += 1

				}
			}
		}
		results[snake.ID] = float64(int(float64(sum)/float64(board.Width*board.Height)*1000)) / 1000
	}
	return results
}

func cellHasFood(board *GameBoard, point Point) bool {
	for _, food := range board.Food {
		if Equals(food, point) {
			return true
		}
	}
	return false
}

func enqueueNextPoints(board *GameBoard, pvalue qNode, hasVisted func(p Point) bool, qPoint *circularbuffer.Queue, visited map[Point]bool) {

	nextDist := pvalue.dist + 1
	//up
	nextPoint := pvalue.p.Clone()
	nextPoint.Y += 1
	if inRange(0, board.Height, nextPoint.Y) && !hasVisted(nextPoint) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint] = true
	}
	//down
	nextPoint = pvalue.p.Clone()
	nextPoint.Y -= 1
	if inRange(0, board.Height, nextPoint.Y) && !hasVisted(nextPoint) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint] = true
	}
	// left
	nextPoint = pvalue.p.Clone()
	nextPoint.X -= 1
	if inRange(0, board.Width, nextPoint.X) && !hasVisted(nextPoint) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint] = true
	}
	// right
	nextPoint = pvalue.p.Clone()
	nextPoint.X += 1
	if inRange(0, board.Width, nextPoint.X) && !hasVisted(nextPoint) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint] = true
	}
}

func (v VoronoiEval) EvaluateBoard(board *GameBoard, snakeId SnakeID, complete bool, count int) float64 {
	scores := VoronoiScore(board)
	res := scores[snakeId]

	healthscore := getHealthScore(&board.Snakes[snakeId])
	if healthscore == 0 {
		return 0
	}
	//s := lengthScore(board, snakeId)
	score := healthscore*0.5 + res*0.5
	if !complete {
		score = score * 0.01 * (float64(count) + 1)
	}
	return score

}
