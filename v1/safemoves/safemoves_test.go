package safemoves

import (
	"jacksnake/models"
	. "jacksnake/models"
	"testing"
)

func Test_getSafeMovesBySnake(t *testing.T) {
	/*
		3 - 0 - 1
		2 - 0 1 1
		1 - - 1 -
		0 - - - -
			0 1 2 3
	*/
	snakes := []Battlesnake{
		{Head: Coord{X: 1, Y: 2},
			Body: []Coord{{X: 1, Y: 2}, {X: 1, Y: 3}}},
		{Head: Coord{X: 3, Y: 3},
			Body: []Coord{{X: 3, Y: 3}, {X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}}}}

	state := GameState{
		You: snakes[0],
		Board: Board{
			Snakes: []Battlesnake{
				snakes[1],
			},
			Width:  4,
			Height: 4,
		},
	}

	res := GetSafeMovesBySnake(state, snakes[1])
	for i := range res {
		print(res[i], " ")
	}
	println()
	if len(res) != 1 {
		println("snake 1 should only have 1 option")
		t.FailNow()
	}

}

func Test_realWorldSafeMoves(t *testing.T) {
	/*
		6 - 0
		5 0 0
		4 0
		3 0
		2
		1
		0 -
			0 1 2 3 4 5 6

	*/
	you := models.Battlesnake{ID: "31bc1af3-5d42-46db-9f5c-5d98dcd8159d", Name: "Go Starter Project", Health: 98,
		Body: []models.Coord{
			models.Coord{X: 1, Y: 6},
			models.Coord{X: 1, Y: 5},
			models.Coord{X: 0, Y: 5},
			models.Coord{X: 0, Y: 4},
			models.Coord{X: 0, Y: 3}},
		Head:   models.Coord{X: 1, Y: 6},
		Length: 5, Latency: "4", Shout: "", Customizations: models.Customizations{Color: "#b13859", Head: "default", Tail: "default"}}

	data := models.GameState{
		Game: models.Game{
			ID: "c64342be-da2e-4c49-95f2-d4f1f47e719d",
			Ruleset: models.Ruleset{Name: "solo", Version: "cli",
				Settings: models.RulesetSettings{FoodSpawnChance: 15, MinimumFood: 1, HazardDamagePerTurn: 14}},
			Map: "standard", Source: "", Timeout: 500},
		Turn: 27,
		Board: models.Board{Height: 7, Width: 7, Food: []models.Coord{
			models.Coord{X: 3, Y: 3},
			models.Coord{X: 3, Y: 6},
			models.Coord{X: 5, Y: 3},
			models.Coord{X: 4, Y: 4}, models.Coord{X: 5, Y: 4}},
			Hazards: []models.Coord{},
			Snakes:  []models.Battlesnake{you}}, You: you}
	moves := GetSafeMovesBySnake(data, you)
	if len(moves) != 2 {
		println("bad number of moves")
		t.FailNow()
	}

}
