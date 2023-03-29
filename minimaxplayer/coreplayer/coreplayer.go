package coreplayer

// models

type Point struct {
	X uint8
	Y uint8
}

func (this *Point) Clone() Point {
	return Point{
		X: this.X,
		Y: this.Y,
	}
}

func Equals(a Point, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}

type Snake struct {
	ID     SnakeID // 0 -4
	Body   []Point
	Health uint8 // 0 - 100
}

func (this *Snake) clone() Snake {
	newBody := make([]Point, len(this.Body))
	for i := range this.Body {
		newBody[i] = this.Body[i].Clone()
	}
	return Snake{
		ID:     this.ID,
		Body:   newBody,
		Health: this.Health,
	}
}

type GameBoard struct {
	Turn    int
	Height  uint8
	Width   uint8
	Food    []Point
	Snakes  []Snake
	Hazards []Point
}

func (this *GameBoard) Clone() GameBoard {
	food := make([]Point, len(this.Food))
	for i := range food {
		food[i] = this.Food[i].Clone()
	}
	hazards := make([]Point, len(this.Hazards))
	for i := range hazards {
		hazards[i] = this.Hazards[i].Clone()
	}
	snakes := make([]Snake, len(this.Snakes))

	for i := range this.Snakes {
		snakes[i] = this.Snakes[i].clone()
	}

	return GameBoard{
		Turn:    this.Turn,
		Height:  this.Height,
		Width:   this.Width,
		Hazards: hazards,
		Snakes:  snakes,
		Food:    food,
	}
}

type SnakeMove struct {
	ID   SnakeID
	Move Direction
}

type (
	Direction int
	SnakeID   uint8
	Score     float32
)

const (
	LEFT Direction = iota
	UP
	RIGHT
	DOWN
)

func DirectionToString(dir Direction) string {
	switch dir {
	case LEFT:
		return "left"
	case UP:
		return "up"
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	}
	println("bad direction received")
	return "error"
}

func StringToDirection(s string) Direction {
	switch s {
	case "left":
		return LEFT
	case "up":
		return UP
	case "right":
		return RIGHT
	case "down":
		return DOWN
	}
	println("bad direction string")
	return DOWN
}

type Simulator interface {
	SimulateMoves(board *GameBoard, moves []SnakeMove)
	GetValidMoves(board *GameBoard, snakeId SnakeID) []Direction
}

type Evaluator interface {
	// evaluates board for a given snake, avaluation should be between 0 and 1
	EvaluateBoard(board *GameBoard, snakeId SnakeID) float64
}

type Player interface {
	Move(board GameBoard, youId SnakeID) Direction
}

type MovingAlgo interface {
	Move(board *GameBoard, youId SnakeID) Direction
}

type PlayerImpl struct {
	simulator  Simulator
	movingAlgo MovingAlgo
}

func (player *PlayerImpl) init(simulator Simulator, moving MovingAlgo) {
	player.simulator = simulator
	player.movingAlgo = moving
}

func (player *PlayerImpl) Move(board GameBoard, youId SnakeID) Direction {
	moves := player.simulator.GetValidMoves(&board, youId)
	if len(moves) == 0 {
		println("No valid moves, moving down")
		return DOWN
	}
	algo := player.movingAlgo
	move := algo.Move(&board, youId)
	return move
}
