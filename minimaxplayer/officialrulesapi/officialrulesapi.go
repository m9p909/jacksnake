package officialrulesapi

import (
	"github.com/BattlesnakeOfficial/rules"
	"jacksnake/minimaxplayer/coreplayer"
)

type OfficialRules interface {
	SimulateMove(board rules.BoardState, move string, snakeID string) rules.BoardState
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
	println("board pointer", &newBoard)
	return adapter.rules.GetValidMoves(newBoard, id)
}
func (adapter *OfficialRulesAdapterImpl) SimulateMove(board coreplayer.GameBoard, move string, snakeId string) coreplayer.GameBoard {
	// stub
	board1 := adapter.converter.ConvertToOfficialBoard(board)
	board2 := adapter.rules.SimulateMove(board1, move, snakeId)
	board3 := adapter.converter.ConvertToCoreBoard(board2)
	return board3
}
