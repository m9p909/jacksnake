package officialrulesapi_test

import (
	"github.com/BattlesnakeOfficial/rules"
	. "jacksnake/minimaxplayer/officialrulesapi"
	"testing"
)

func parametrizedGetValidMoves(t *testing.T, state rules.BoardState, id string, moves map[string]int) {
	rules := OfficialRulesImpl{}
	res := rules.GetValidMoves(state, id)

	if len(res) != len(moves) {
		t.Errorf("Expected %d moves, got %d", len(moves), len(res))
	}

	for _, move := range res {
		if _, ok := moves[move]; !ok {
			t.Errorf("Unexpected move: %s", move)
		}
	}

}

func Test_GetValidMoves(t *testing.T) {
	parametrizedGetValidMoves(t, getRulesState(), "1", map[string]int{"down": 1, "left": 1})
	parametrizedGetValidMoves(t, getRulesState(), "2", map[string]int{"up": 1, "down": 1})

}
