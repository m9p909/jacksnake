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

type Simulator interface {
	SimulateMove(board GameBoard, move string) GameBoard
	GetValidMoves(board GameBoard, move string) []string
}

type Player interface {
	Move(board GameBoard, youId string) string
}
