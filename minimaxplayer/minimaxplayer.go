package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type MinimaxPlayer struct {
}

func (player *MinimaxPlayer) Move(state models.GameState) string {
	core :=
		StateConverter.StateToCore(state)
	move := coreplayer.Player.Move(core, state.You.ID)
	return move
}

func (player *MinimaxPlayer) Start(state models.GameState) {
	println("START MINIMAX")
}
func (player *MinimaxPlayer) End(state models.GameState) {
	println("END MINIMAX")
}

	StateToCore(state models.GameState) coreplayer.GameBoard
	CoreToState(state models.GameState) coreplayer.GameBoard
