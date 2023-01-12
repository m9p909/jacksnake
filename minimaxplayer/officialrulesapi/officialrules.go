package officialrulesapi

import (
	"github.com/BattlesnakeOfficial/rules"
)

var standardRulesetStages = []string{
	rules.StageGameOverStandard,
	rules.StageMovementStandard,
	rules.StageStarvationStandard,
	rules.StageHazardDamageStandard,
	rules.StageFeedSnakesStandard,
	rules.StageEliminationStandard,
}

type OfficialRulesImpl struct {
}

var standardRules = rules.NewSettingsWithParams(standardRulesetStages...)

func (*OfficialRulesImpl) rulesSimulateMove(board rules.BoardState, move string, snakeID string) (bool, rules.BoardState, error) {
	success, err := rules.MoveSnakesStandard(&board, standardRules, []rules.SnakeMove{{ID: snakeID, Move: move}})
	return success, board, err
}

func (officialRules *OfficialRulesImpl) SimulateMove(board rules.BoardState, move string, snakeID string) rules.BoardState {
	_, board, err := officialRules.rulesSimulateMove(board, move, snakeID)

	if err != nil {
		println("could not simulate move")
	}

	return board
}

var moves = []string{"up", "down", "left", "right"}

func (officialRules *OfficialRulesImpl) GetValidMoves(board rules.BoardState, snakeID string) []string {
	moves := []string{}
	for _, move := range moves {

		success, _, err := officialRules.rulesSimulateMove(board, move, snakeID)
		if err != nil {
			println("could not GetValidMoves")
		}
		if success {
			moves = append(moves, move)
		}
	}

	return moves
}
