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
	moves := adapter.converter.ConvertToOfficialBoard(board)
	return adapter.rules.GetValidMoves(moves, id)
}
func (*OfficialRulesAdapterImpl) SimulateMove(board coreplayer.GameBoard, move string, snakeId string) coreplayer.GameBoard {
	// stub
	return coreplayer.GameBoard{}
}
