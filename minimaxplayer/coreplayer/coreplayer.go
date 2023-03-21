package coreplayer

// models

type Point struct {
	X int
	Y int
}

func Equals(a Point, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}

type Snake struct {
	ID     string
	Body   []Point
	Health int
}

type GameBoard struct {
	Turn    int
	Height  int
	Width   int
	Food    []Point
	Snakes  []Snake
	Hazards []Point
}

type SnakeMove struct {
	ID   string
	Move string
}

const (
	LEFT  = 0
	UP    = 1
	RIGHT = 2
	DOWN  = 3
)

type Simulator interface {
	SimulateMoves(board GameBoard, moves []SnakeMove) GameBoard
	GetValidMoves(board GameBoard, snakeId string) []string
}

type Evaluator interface {
	// evaluates board for a given snake, avaluation should be between 0 and 1
	EvaluateBoard(board *GameBoard, snakeId string) float64
}

type Player interface {
	Move(board GameBoard, youId string) string
}

type MovingAlgo interface {
	Move(board GameBoard, youId string) string
}

type PlayerImpl struct {
	simulator  Simulator
	movingAlgo MovingAlgo
}

func (player *PlayerImpl) init(simulator Simulator, moving MovingAlgo) {
	player.simulator = simulator
	player.movingAlgo = moving
}

func (player *PlayerImpl) Move(board GameBoard, youId string) string {
	moves := player.simulator.GetValidMoves(board, youId)
	if len(moves) == 0 {
		println("No valid moves, moving down")
		return "down"
	}
	algo := player.movingAlgo
	move := algo.Move(board, youId)
	return move
}
