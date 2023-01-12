package officialrulesapi_test

import (
	"github.com/BattlesnakeOfficial/rules"
	"testing"
)

/*
	type OfficialRulesAdapter interface {
		SimulateMove(board coreplayer.GameBoard, move string, snakeId string) coreplayer.GameBoard
		GetValidMoves(board coreplayer.GameBoard, snakeID string) []string
	}
*/
func Test_OfficialRulesAdapter(t *testing.T) {

}

var standardRulesetStages = []string{
	rules.StageGameOverStandard,
	rules.StageMovementStandard,
	rules.StageStarvationStandard,
	rules.StageHazardDamageStandard,
	rules.StageFeedSnakesStandard,
	rules.StageEliminationStandard,
}
