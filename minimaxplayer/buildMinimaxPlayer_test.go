package minimaxplayer_test

import (
	"jacksnake/minimaxplayer"
	"testing"
)

func Test_randomPlayer(t *testing.T) {
	player := minimaxplayer.BuildRandomPlayer()
	move := player.Move(getGameStateTest1())
	if move == "right" {
		println("Picked a bad move, can't go right")
		t.Fail()
	}
}
