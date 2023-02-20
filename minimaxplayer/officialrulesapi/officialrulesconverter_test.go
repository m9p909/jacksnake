package officialrulesapi_test

import (
	"encoding/json"
	"jacksnake/minimaxplayer/coreplayer"
	. "jacksnake/minimaxplayer/officialrulesapi"
	"testing"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/gkampitakis/go-snaps/snaps"
)

func getCoreState() coreplayer.GameBoard {
	/*
		3 - 0 - f
		2 - 0 1 1
		1 x - 1 f
		0 - x - -
			0 1 2 3
	*/

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
		Food:    food,
		Hazards: hazards,
		Snakes:  snakes,
	}
}

func getRulesState() rules.BoardState {
	/*
		3 - 0 - f
		2 - 0 1 1
		1 x - 1 f
		0 - x - -
			0 1 2 3
	*/
	food := []rules.Point{
		{X: 3, Y: 1},
		{X: 3, Y: 3},
	}

	hazards := []rules.Point{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	snakes := []rules.Snake{
		{
			ID:     "1",
			Body:   []rules.Point{{X: 1, Y: 2}, {X: 1, Y: 3}},
			Health: 99,
		},
		{
			ID:     "2",
			Body:   []rules.Point{{X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}},
			Health: 100,
		},
	}

	return rules.BoardState{
		Turn:    2,
		Height:  5,
		Width:   4,
		Food:    food,
		Hazards: hazards,
		Snakes:  snakes,
	}

}

func Test_CoretoRules(t *testing.T) {
	converter := OfficialRulesConverter{}

	board := converter.ConvertToOfficialBoard(getCoreState())

	if board.Turn != 2 || board.Height != 5 || board.Width != 4 {
		println("bad board vaue")
		println(json.Marshal(board))
	}

	snaps.MatchSnapshot(t, board.Food)
	snaps.MatchSnapshot(t, board.Hazards)
	snaps.MatchSnapshot(t, board.Snakes)

}

func Test_RulestoCore(t *testing.T) {
	converter := OfficialRulesConverter{}

	board := converter.ConvertToOfficialBoard(getCoreState())

	if board.Turn != 2 || board.Height != 5 || board.Width != 4 {
		println("bad board vaue")
		println(json.Marshal(board))
	}

	snaps.MatchSnapshot(t, board.Food)
	snaps.MatchSnapshot(t, board.Hazards)
	snaps.MatchSnapshot(t, board.Snakes)

}

func TestSnakeMovesToCore(t *testing.T) {
	converter := OfficialRulesConverter{}

	move := []rules.SnakeMove{
		{
			ID:   "1",
			Move: "up",
		},
		{
			ID:   "2",
			Move: "down",
		},
	}

	moveSnake := converter.ConvertSnakeMovesToCore(move)

	snaps.MatchSnapshot(t, moveSnake)

}
