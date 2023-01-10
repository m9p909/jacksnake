package officialrulesapi

type OfficialRulesAdapter interface {
	convertToOfficialBoard(state coreplayer.GameBoard) rules.BoardState
	convertFromOfficialBoard(state rules.BoardState) coreplayer.GameBoard
	SimulateMove(board coreplayer.GameBoard, move string, snakeID string) coreplayer.GameBoard
	GetValidMoves(board coreplayer.GameBoard, snakeID string) []string
}

type OfficialRules interface {
	SimulateMove(board rules.GameBoard, move string, snakeID string) rules.GameBoard
	GetValidMoves(board rules.GameBoard, snakeID string) []string
}




type OfficialRulesAdapterImpl struct {
	rules OfficialRules
}

func (adapter *OfficialRulesAdapter) init(rules OfficialRules) {
	adapter.rules = rules
}

func (*OfficialRulesAdapter) convertToOfficialBoard(state coreplayer.GameBoard) rules.BoardState{

}

func (adapter *OfficialRulesAdapter) GetValidMoves(board coreplayer.GameBoard) []string {

}

func (adapter *OfficialRulesAdapter) SimulateMove(coreplayer.GameBoard, move string, snakeId string) coreplayer.GameBoard {
	// stub
	return coreplayer.GameBoard{}
}
