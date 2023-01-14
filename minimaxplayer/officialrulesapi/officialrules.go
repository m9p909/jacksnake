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
	if err != nil {
		println("something went wrong")
	}
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

var movesConst = []string{"up", "down", "left", "right"}

func (officialRules *OfficialRulesImpl) GetValidMoves(board rules.BoardState, snakeID string) []string {
	output := []string{}
	if board.Snakes == nil || len(board.Snakes) == 0 {
		println("snakes nil")
		return []string{"down"}
	}
	// board is null
	for _, move := range movesConst {

		testBoard := board.Clone()
		success, _, err := officialRules.rulesSimulateMove(*testBoard, move, snakeID)
		if err != nil {
			println("could not GetValidMoves")
		}
		if success {
			output = append(output, move)
		}
	}

	return output
}
