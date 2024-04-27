package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/customsimulator"
	"jacksnake/minimaxplayer/evaluator"
	"jacksnake/models"
)

type PerRequestPlayer struct {
	converter StateConverter
}

func (player *PerRequestPlayer) Init(conv StateConverter) {
	player.converter = conv
}

func (player *PerRequestPlayer) Move(state models.GameState) string {
	// printJson(state)
	// clone the player
	core, id := player.converter.StateToCore(state)
	total := 0
	for _, snake := range core.Snakes {
		if snake.Health > 0 {
			total++
		}
	}
	algo := coreplayer.NewMinimaxAlgoMove(customsimulator.New(), evaluator.NewVoronoiEval(), 10-total)
	move := algo.Move(core, id)
	return coreplayer.DirectionToString(move)
}

func (player *PerRequestPlayer) Start(state models.GameState) {
	println("START MINIMAX")
}

func (player *PerRequestPlayer) End(state models.GameState) {
	println("END MINIMAX")
}
