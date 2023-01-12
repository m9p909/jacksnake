package officialrulesapi_test

import (
	"encoding/json"
	"jacksnake/minimaxplayer/coreplayer"
	. "jacksnake/minimaxplayer/officialrulesapi"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func getGameState1Result() coreplayer.GameBoard {
	food := []coreplayer.Point{
		{X: 3, Y: 1},
		{X: 3, Y: 3},
	}

	hazards := []coreplayer.Point{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	snakes := []coreplayer.Snake{
		{
			ID:     "1",
			Body:   []coreplayer.Point{{X: 1, Y: 2}, {X: 1, Y: 3}},
			Health: 99,
		},
		{
			ID:     "2",
			Body:   []coreplayer.Point{{X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}},
			Health: 100,
		},
	}

	return coreplayer.GameBoard{
		Turn:    2,
		Height:  5,
		Width:   4,
		Food:    []coreplayer.Point{},
		Hazards: hazards,
		Snakes:  snakes,
	}
}

func Test_OfficialRulesConverter(t *testing.T) {
	converter := OfficialRulesConverter{}

	board := converter.ConvertToOfficialBoard(getGameState1Result())

	if board.Turn != 2 || board.Height != 5 || board.Width != 4 {
		println("bad board vaue")
		println(json.Marshal(board))
	}

	snaps.MatchSnapshot(t, board.Food)
	snaps.MatchSnapshot(t, board.Hazards)
	snaps.MatchSnapshot(t, board.Snakes)

}
