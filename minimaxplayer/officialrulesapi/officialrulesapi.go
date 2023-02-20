package officialrulesapi

import (
	"github.com/BattlesnakeOfficial/rules"
	"jacksnake/minimaxplayer/coreplayer"
)

type OfficialRules interface {
	SimulateMoves(board rules.BoardState, moves []rules.SnakeMove) rules.BoardState
	GetValidMoves(board rules.BoardState, snakeID string) []string
}

type OfficialRulesAdapterImpl struct {
	rules     OfficialRules
	converter OfficialRulesConverter
}

func (adapter *OfficialRulesAdapterImpl) init(rules OfficialRules) {
	adapter.rules = rules
	adapter.converter = OfficialRulesConverter{}
}

func (adapter *OfficialRulesAdapterImpl) GetValidMoves(board coreplayer.GameBoard, id string) []string {
	newBoard := adapter.converter.ConvertToOfficialBoard(board)
	return adapter.rules.GetValidMoves(newBoard, id)
}

func (adapter *OfficialRulesAdapterImpl) SimulateMoves(board coreplayer.GameBoard, snakeMoves []coreplayer.SnakeMove) coreplayer.GameBoard {
	board1 := adapter.converter.ConvertToOfficialBoard(board)
	moves := adapter.converter.ConvertSnakeMovesToRules(snakeMoves)
	board2 := adapter.rules.SimulateMoves(board1, moves)
	board3 := adapter.converter.ConvertToCoreBoard(board2)
	return board3
}
