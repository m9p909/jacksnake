package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type StateConverterImpl struct{}

func (*StateConverterImpl) coordArrToPointArr(coords []models.Coord) []coreplayer.Point {
	points := []coreplayer.Point{}

	for _, coord := range coords {
		p := coreplayer.Point{
			X: uint8(coord.X),
			Y: uint8(coord.Y),
		}
		points = append(points, p)
	}
	return points
}

func removeDuplicates(arr []coreplayer.Point) []coreplayer.Point {
	keys := make(map[coreplayer.Point]bool)
	list := []coreplayer.Point{}
	for _, item := range arr {
		if _, value := keys[item]; !value {
			keys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func (conv *StateConverterImpl) battleSnakeToSnake(snake []models.Battlesnake, youId string) ([]coreplayer.Snake, coreplayer.SnakeID) {
	snakes := []coreplayer.Snake{}

	id := 0
	outId := 0
	for _, snake := range snake {
		nextSnake := coreplayer.Snake{
			ID:     coreplayer.SnakeID(id),
			Health: uint8(snake.Health),
			Body:   conv.coordArrToPointArr(snake.Body),
		}
		nextSnake.Body = removeDuplicates(nextSnake.Body)
		snakes = append(snakes, nextSnake)
		if snake.ID == youId {
			outId = id
		}
		id++
	}

	return snakes, coreplayer.SnakeID(outId)
}

func (conv *StateConverterImpl) StateToCore(state models.GameState) (coreplayer.GameBoard, coreplayer.SnakeID) {
	snakes, id := conv.battleSnakeToSnake(state.Board.Snakes, state.You.ID)
	board := coreplayer.GameBoard{
		Turn:    state.Turn,
		Height:  uint8(state.Board.Height),
		Width:   uint8(state.Board.Width),
		Food:    conv.coordArrToPointArr(state.Board.Food),
		Hazards: conv.coordArrToPointArr(state.Board.Hazards),
		Snakes:  snakes,
	}
	return board, id
}
