package officialrulesapi_test

import (
	"github.com/BattlesnakeOfficial/rules"
	"testing"
)

/*

 */

var standardRulesetStages = []string{
	rules.StageGameOverStandard,
	rules.StageMovementStandard,
	rules.StageStarvationStandard,
	rules.StageHazardDamageStandard,
	rules.StageFeedSnakesStandard,
	rules.StageEliminationStandard,
}

func Test_createsABoard(t *testing.T) {
	state := rules.BoardState{}
	// uses board state, only uses the snake Body, does not use the map
	rules.MoveSnakesStandard()
	rules.NewPipeline(standardRulesetStages...)
}
