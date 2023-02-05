package officialrulesapi

import (
	"errors"
	"github.com/BattlesnakeOfficial/rules"
	"jacksnake/minimaxplayer/officialrulesapi/safemoves"
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

func (*OfficialRulesImpl) rulesSimulateMove(board rules.BoardState,
	snakeMoves []rules.SnakeMove) (rules.BoardState, error) {

	_, err := rules.MoveSnakesStandard(&board, standardRules, snakeMoves)
	return board, err
}

func (officialRules *OfficialRulesImpl) SimulateMoves(board rules.BoardState, snakeMoves []rules.SnakeMove) rules.BoardState {
	board, err := officialRules.rulesSimulateMove(board, snakeMoves)
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
