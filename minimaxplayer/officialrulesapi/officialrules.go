package officialrulesapi

import (
	"errors"

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

func (*OfficialRulesImpl) rulesSimulateMove(board rules.BoardState, move string, snakeID string) (bool, rules.BoardState, error) {
	rules.MoveSnakesStandard(&board, standardRules, []rules.SnakeMove{{ID: snakeID, Move: move}})
	snake, err := getSnake(board, snakeID)
	success := snake.EliminatedCause == ""
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
