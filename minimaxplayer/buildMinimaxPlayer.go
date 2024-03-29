package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/customsimulator"
	"jacksnake/minimaxplayer/evaluator"
	. "jacksnake/models"
)

type Player interface {
	Move(state GameState) string // string is 1 of up down left or right
	Start(state GameState)
	End(state GameState)
}

func BuildRandomPlayer() Player {
	player := MinimaxPlayer{}
	conv := StateConverterImpl{}
	simulator := customsimulator.New()
	algo := coreplayer.NewRandomAlgo(simulator)
	player.Init(&conv, algo)
	return &player
}

func BuildMinimaxPlayer() Player {
	conv := StateConverterImpl{}
	algo := coreplayer.NewMinimaxAlgoMove(customsimulator.New(),
		evaluator.NewAgroEvaluator(), 6)
	player := MinimaxPlayer{}
	player.Init(&conv, algo)
	return &player
}

func BuildHeuristicPlayer() Player {
	conv := StateConverterImpl{}
	algo := coreplayer.NewHeuristicAlgo(customsimulator.New(),
		evaluator.NewAgroEvaluator())
	player := MinimaxPlayer{}
	player.Init(&conv, algo)
	return &player
}
