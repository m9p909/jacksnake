package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"fmt"
	"jacksnake/minimaxplayer"
	. "jacksnake/models"
	"log"
	"time"
)

type Player interface {
	Move(state GameState) string // string is 1 of up down left or right
	Start(state GameState)
	End(state GameState)
}

type MockPlayer struct{}

func (*MockPlayer) Move(_ GameState) string {
	return "down"
}

func (*MockPlayer) Start(_ GameState) {}

func (*MockPlayer) End(_ GameState) {}

type MainResponder struct {
	player Player
}

func (res *MainResponder) init(player Player) {
	res.player = player
}

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func (*MainResponder) Info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "m9p909",   // TODO: Your Battlesnake username
		Color:      "#b13859",  // TODO: Choose color
		Head:       "dead",     // TODO: Choose head
		Tail:       "mlh-gene", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func (res *MainResponder) Start(state GameState) {
	log.Println("GAME START")
	res.player.Start(state)
}

// end is called when your Battlesnake finishes a game
func (responder *MainResponder) End(state GameState) {
	log.Printf("GAME OVER\n\n")
	responder.player.End(state)
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func (responder *MainResponder) Move(state GameState) BattlesnakeMoveResponse {
	t1 := time.Now()

	nextMove := responder.player.Move(state)

	fmt.Printf("time: %s\n", time.Since(t1))

	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	res := MainResponder{}
	player := minimaxplayer.BuildMinimaxPlayer()
	res.init(player)
	RunServer(&res)
}
