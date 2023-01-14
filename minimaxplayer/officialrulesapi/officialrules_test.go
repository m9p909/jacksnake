package officialrulesapi_test

import (
	"github.com/gkampitakis/go-snaps/snaps"
	. "jacksnake/minimaxplayer/officialrulesapi"
	"testing"
)

func Test_GetValidMoves(t *testing.T) {
	state := getRulesState()

	rules := OfficialRulesImpl{}
	res := rules.GetValidMoves(state, "1")
	snaps.MatchSnapshot(t, res)
}
