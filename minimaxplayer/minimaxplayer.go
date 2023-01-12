package minimaxplayer

// gbooch@us.ibm.com

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type StateConverter interface {
	StateToCore(state models.GameState) coreplayer.GameBoard
	CoreToState(state coreplayer.GameBoard) models.GameState
}

type MinimaxPlayer struct {
	converter StateConverter
	player    coreplayer.Player
}

func (player *MinimaxPlayer) Init(conv StateConverter, coreplayer coreplayer.Player) {
	player.converter = conv
	player.player = coreplayer
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
