package evaluator

import (
	"github.com/emirpasic/gods/queues/circularbuffer"
	. "jacksnake/minimaxplayer/coreplayer"
	"sync"
)

type BoardCache struct {
	m    map[*GameBoard][]float64
	lock sync.RWMutex
}

func (c *BoardCache) read(v *GameBoard) ([]float64, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	val, foo := c.m[v]
	return val, foo
}

func (c *BoardCache) write(v *GameBoard, float642 []float64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.m[v] = float642
}

func NewBordCache() *BoardCache {
	return &BoardCache{m: make(map[*GameBoard][]float64)}
}

type VoronoiEval struct {
	cache *BoardCache
}

func NewVoronoiEval() Evaluator {
	return &VoronoiEval{cache: NewBordCache()}
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

func VoronoiScore(board *GameBoard, cache *BoardCache) []float64 {
	val, ok := cache.read(board)
	if ok {
		return val
	}
	arrBoard := buildVoronoiSnakeBoard(board.Snakes, board.Height, board.Width)
	for _, snake := range board.Snakes {
		if snake.Health <= 0 {
			continue
		}
		visited := makeEmptyBoard(board.Height, board.Width)
		head := snake.Body[0]
		qPoint := circularbuffer.New(int(board.Height+board.Width) * 2)
		enqueueNextPoints(board, qNode{head, 0}, qPoint, visited)
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

			enqueueNextPoints(board, pvalue, qPoint, visited)
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
	cache.write(board, results)
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

func enqueueNextPoints(board *GameBoard, pvalue qNode, qPoint *circularbuffer.Queue, visited [][]uint8) {

	nextDist := pvalue.dist + 1
	//up
	nextPoint := pvalue.p.Clone()
	nextPoint.Y += 1
	if inRange(0, board.Height, nextPoint.Y) && (visited[nextPoint.Y][nextPoint.X] == 0) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint.Y][nextPoint.X] = 1
	}
	//down
	nextPoint = pvalue.p.Clone()
	nextPoint.Y -= 1
	if inRange(0, board.Height, nextPoint.Y) && (visited[nextPoint.Y][nextPoint.X] == 0) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint.Y][nextPoint.X] = 1
	}
	// left
	nextPoint = pvalue.p.Clone()
	nextPoint.X -= 1
	if inRange(0, board.Width, nextPoint.X) && (visited[nextPoint.Y][nextPoint.X] == 0) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint.Y][nextPoint.X] = 1
	}
	// right
	nextPoint = pvalue.p.Clone()
	nextPoint.X += 1
	if inRange(0, board.Width, nextPoint.X) && (visited[nextPoint.Y][nextPoint.X] == 0) {
		qPoint.Enqueue(qNode{nextPoint, nextDist})
		visited[nextPoint.Y][nextPoint.X] = 1
	}
}

func (v *VoronoiEval) EvaluateBoard(board *GameBoard, snakeId SnakeID, complete bool, count int) float64 {
	healthscore := getHealthScore(&board.Snakes[snakeId])
	if healthscore == 0 {
		return 0
	}

	scores := VoronoiScore(board, v.cache)
	res := scores[snakeId]

	//s := lengthScore(board, snakeId)
	score := healthscore*0.5 + res*0.5
	if !complete {
		score = score * 0.01 * (float64(count) + 1)
	}
	return score

}
