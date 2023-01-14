package officialrulesapi

import (
	"errors"
	"jacksnake/minimaxplayer/officialrulesapi/safemoves"

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

func getSnake(board rules.BoardState, snakeID string) (*rules.Snake, error) {
	for _, snake := range board.Snakes {
		if snake.ID == snakeID {
			return &snake, nil
		}
	}
	return nil, errors.New("cant find snake")
}

func (*OfficialRulesImpl) rulesSimulateMove(board rules.BoardState, move string, snakeID string) (rules.BoardState, error) {
	_, err := rules.MoveSnakesStandard(&board, standardRules, []rules.SnakeMove{{ID: snakeID, Move: move}})
	return board, err
}

func (officialRules *OfficialRulesImpl) SimulateMove(board rules.BoardState, move string, snakeID string) rules.BoardState {
	board, err := officialRules.rulesSimulateMove(board, move, snakeID)

	if err != nil {
		println("could not simulate move")
	}

	return board
}

var movesConst = []string{"up", "down", "left", "right"}

func (rules *OfficialRulesImpl) snakeIsDead(board *rules.BoardState, snakeID string) bool {
	snek, _ := getSnake(*board, snakeID)
	return snek.EliminatedBy != ""

}

func (officialRules *OfficialRulesImpl) GetValidMoves(board rules.BoardState, snakeID string) []string {
	return safemoves.GetSafeMovesBySnake(board, snakeID)
}
