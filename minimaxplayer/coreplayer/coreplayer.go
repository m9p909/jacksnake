package coreplayer

// models

type Point struct {
	X int
	Y int
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

type Simulator interface {
	SimulateMoves(board GameBoard, moves []SnakeMove) GameBoard
	GetValidMoves(board GameBoard, snakeId string) []string
}

type Evaluator interface {
	// evaluates board for a given snake, avaluation should be between 0 and 1
	EvaluateBoard(board GameBoard, snakeId string)
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
	move := player.movingAlgo.Move(board, youId)
	return move
}
