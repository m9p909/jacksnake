package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/officialrulesapi"
	. "jacksnake/models"
)

type Player interface {
	Move(state GameState) string // string is 1 of up down left or right
	Start(state GameState)
	End(state GameState)
}

func BuildMinimaxPlayer() Player {
	player := MinimaxPlayer{}
	conv := StateConverterImpl{}
	algo := coreplayer.RandomAlgo{}
	standardRulesSimulator := officialrulesapi.GetOfficialRules()
	algo.Init(standardRulesSimulator)
	player.Init(&conv, &algo)
	return &player
}
