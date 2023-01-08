package officialrulesapi_test

import (
	"github.com/BattlesnakeOfficial/rules"
	"jacksnake/minimaxplayer/coreplayer"
	"testing"
)

/*

 */

type OfficialRulesAdapter interface {
	convertToOfficialBoard(state coreplayer.GameBoard) rules.BoardState
	convertFromOfficialBoard(state rules.BoardState) coreplayer.GameBoard
	SimulateMove(coreplayer.GameBoard, move string) coreplayer.GameBoard
	GetValidMoves(board coreplayer.GameBoard, snakeID string) string[]
}

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
