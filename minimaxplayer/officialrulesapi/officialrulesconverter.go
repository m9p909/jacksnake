package officialrulesapi

import (
	"github.com/BattlesnakeOfficial/rules"
	"jacksnake/minimaxplayer/coreplayer"
)

type OfficialRulesConverter struct{}

func (*OfficialRulesConverter) convertRulesPointToCorePoint(point rules.Point) coreplayer.Point {
	return coreplayer.Point{
		X: point.X,
		Y: point.Y,
	}
}

func (adapter *OfficialRulesConverter) convertCorePointToRulesPoint(point coreplayer.Point) rules.Point {
	return rules.Point{
		X: point.X,
		Y: point.Y,
	}
}

func (adapter *OfficialRulesConverter) convertCorePointsToRulesPoints(point []coreplayer.Point) []rules.Point {
	rulesPoints := make([]rules.Point, len(point))
	for i, p := range point {
		rulesPoints[i] = adapter.convertCorePointToRulesPoint(p)
	}
	return rulesPoints
}

func (adapter *OfficialRulesConverter) convertRulesPointsToCorePoints(point []rules.Point) []coreplayer.Point {
	corePoints := make([]coreplayer.Point, len(point))
	for i, p := range point {
		corePoints[i] = adapter.convertRulesPointToCorePoint(p)
	}
	return corePoints
}

func (adaper *OfficialRulesConverter) convertBoardSnakeToRulesSnake(snake []coreplayer.Snake) []rules.Snake {
	rulesSnakes := make([]rules.Snake, len(snake))
	for i, s := range snake {
		rulesSnakes[i] = rules.Snake{
			ID:     s.ID,
			Health: s.Health,
			Body:   adaper.convertCorePointsToRulesPoints(s.Body),
		}
	}
	return rulesSnakes
}

func (adapter *OfficialRulesConverter) ConvertToOfficialBoard(state coreplayer.GameBoard) rules.BoardState {
	newState := rules.BoardState{
		Turn:    state.Turn,
		Height:  state.Height,
		Width:   state.Width,
		Food:    adapter.convertCorePointsToRulesPoints(state.Food),
		Hazards: adapter.convertCorePointsToRulesPoints(state.Hazards),
		Snakes:  adapter.convertBoardSnakeToRulesSnake(state.Snakes),
	}
	return newState
}
