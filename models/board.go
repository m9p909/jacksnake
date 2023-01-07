package models

const (
	hazardSymbol = "x"
	foodSymbol   = "f"
)

type Grid struct {
	board  [][]string
	snakes []Battlesnake
}

func NewGrid(state GameState) {

}

func (grid *Grid) SetHazard(coord Coord) {

}

func (grid *Grid) GetHazard(coord Coord) {

}
