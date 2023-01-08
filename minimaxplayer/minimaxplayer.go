package minimaxplayer

import (
	"jacksnake/models"
)

type MinimaxPlayer struct {
}

func (player *MinimaxPlayer) Move(state models.GameState) string {
	return "down"
}

func (player *MinimaxPlayer) Start(state models.GameState) {}
func (player *MinimaxPlayer) End(state models.GameState)   {}
