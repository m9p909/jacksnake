package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type StateConverter interface {
	StateToCore(state models.GameState) coreplayer.GameBoard
	CoreToState(state models.GameState) coreplayer.GameBoard
}

type MinimaxPlayer struct {
	converter StateConverter
	player    coreplayer.Player
}

func (player *MinimaxPlayer) init(conv StateConverter) {
	player.converter = conv
}

func (player *MinimaxPlayer) Move(state models.GameState) string {
	core := player.converter.StateToCore(state)
	move := player.player.Move(core, state.You.ID)
	return move
}

func (player *MinimaxPlayer) Start(state models.GameState) {
	println("START MINIMAX")
}
func (player *MinimaxPlayer) End(state models.GameState) {
	println("END MINIMAX")
}
