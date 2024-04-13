package minimaxplayer

import (
	"encoding/json"
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type StateConverter interface {
	StateToCore(state models.GameState) (coreplayer.GameBoard, coreplayer.SnakeID)
}

type MinimaxPlayer struct {
	converter StateConverter
	player    coreplayer.Player
}

func (player *MinimaxPlayer) Init(conv StateConverter, coreplayer coreplayer.Player) {
	player.converter = conv
	player.player = coreplayer
}

func printJson(state models.GameState) {
	data, _ := json.Marshal(state)
	println(string(data))
}

func (player *MinimaxPlayer) Move(state models.GameState) string {
	// printJson(state)
	// clone the player
	core, id := player.converter.StateToCore(state)
	move := player.player.Move(core, id)
	return coreplayer.DirectionToString(move)
}

func (player *MinimaxPlayer) Start(state models.GameState) {
	println("START MINIMAX")
}

func (player *MinimaxPlayer) End(state models.GameState) {
	println("END MINIMAX")
}
