package officialrulesapi_test

import (
	"jacksnake/minimaxplayer/officialrulesapi"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

/*
	type OfficialRulesAdapter interface {
		SimulateMove(board coreplayer.GameBoard, move string, snakeId string) coreplayer.GameBoard
		GetValidMoves(board coreplayer.GameBoard, snakeID string) []string
	}
*/
func Test_OfficialRulesAdapter(t *testing.T) {
	rules := officialrulesapi.GetOfficialRules()
	moves := rules.GetValidMoves(getCoreState(), "1")
	snaps.MatchSnapshot(t, moves)

}
