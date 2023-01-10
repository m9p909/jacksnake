package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type StateConverterImpl struct{}

func (StateConverterImpl) coordArrToPointArr(coord []models.Coord) []coreplayer.Point {
	points := []coreplayer.Point{}

	for _, coord := range coord {
		points = append(points, coreplayer.Point(coord))
	}
	return points
}

func (StateConverterImpl) pointArrToCoordArr(points []coreplayer.Point) []models.Coord {
	coords := []models.Coord{}

	for _, point := range points {
		coords = append(coords, models.Coord(point))
	}

	return coords
}

func (conv *StateConverterImpl) battleSnakeToSnake(snake []models.Battlesnake) []coreplayer.Snake {
	snakes := []coreplayer.Snake{}

	for _, snake := range snake {
		nextSnake := coreplayer.Snake{
			ID:     snake.ID,
			Health: snake.Health,
			Body:   conv.coordArrToPointArr(snake.Body),
		}
		snakes = append(snakes, nextSnake)
	}

	return snakes

}

func (conv *StateConverterImpl) snakeToBattlesnake(snake []coreplayer.Snake) []models.Battlesnake {
	snakes := []models.Battlesnake{}

	for _, snake := range snake {
		nextSnake := models.Battlesnake{
			ID:     snake.ID,
			Health: snake.Health,
			Body:   conv.pointArrToCoordArr(snake.Body),
		}
		snakes = append(snakes, nextSnake)
	}

	return snakes
}

func (conv *StateConverterImpl) StateToCore(state models.GameState) coreplayer.GameBoard {
	board := coreplayer.GameBoard{
		Turn:    state.Turn,
		Height:  state.Board.Height,
		Width:   state.Board.Width,
		Food:    conv.coordArrToPointArr(state.Board.Food),
		Hazards: conv.coordArrToPointArr(state.Board.Hazards),
		Snakes:  conv.battleSnakeToSnake(state.Board.Snakes),
	}
	return board
}

func (conv *StateConverterImpl) CoreToState(state coreplayer.GameBoard) models.GameState {
	return models.GameState{
		Turn: state.Turn,
		Board: models.Board{
			Height:  state.Height,
			Width:   state.Width,
			Food:    conv.pointArrToCoordArr(state.Food),
			Hazards: conv.pointArrToCoordArr(state.Hazards),
			Snakes:  conv.snakeToBattlesnake(state.Snakes),
		},
	}
}
