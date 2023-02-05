package officialrulesapi

import (
	"github.com/BattlesnakeOfficial/rules"
	"jacksnake/minimaxplayer/coreplayer"
)

// Privacy and Solitude Diana Webb <- cool book

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

func (adapter *OfficialRulesConverter) convertRulesSnakeToBoardSnake(snakes []rules.Snake) []coreplayer.Snake {
	rulesSnakes := make([]coreplayer.Snake, len(snakes))
	for i, snake := range snakes {
		rulesSnakes[i] = coreplayer.Snake{
			ID:     snake.ID,
			Health: snake.Health,
			Body:   adapter.convertRulesPointsToCorePoints(snake.Body),
		}
	}
	return rulesSnakes
}

func (adapter *OfficialRulesConverter) ConvertToCoreBoard(state rules.BoardState) coreplayer.GameBoard {
	newState := coreplayer.GameBoard{
		Turn:    state.Turn,
		Height:  state.Height,
		Width:   state.Width,
		Food:    adapter.convertRulesPointsToCorePoints(state.Food),
		Hazards: adapter.convertRulesPointsToCorePoints(state.Hazards),
		Snakes:  adapter.convertRulesSnakeToBoardSnake(state.Snakes),
	}
	return newState

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

func (adapter *OfficialRulesConverter) ConvertSnakeMovesToRules(moves []coreplayer.SnakeMove) []rules.SnakeMove {
	rulesMoves := make([]rules.SnakeMove, len(moves))
	for i, move := range moves {
		rulesMoves[i] = rules.SnakeMove{
			ID:   move.ID,
			Move: move.Move,
		}
	}
	return rulesMoves
}

// Function that converts rules Snake Move structure to coreplayer snake move structure
func (adapter *OfficialRulesConverter) ConvertSnakeMovesToCore(moves []rules.SnakeMove) []coreplayer.SnakeMove {
	coreMoves := make([]coreplayer.SnakeMove, len(moves))
	for i, move := range moves {
		coreMoves[i] = coreplayer.SnakeMove{
			ID:   move.ID,
			Move: move.Move,
		}
	}
	return coreMoves
}
