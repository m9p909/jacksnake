package minimaxplayer_test

import (
	"encoding/json"
	"fmt"
	"jacksnake/minimaxplayer"
	"jacksnake/models"
	. "jacksnake/models"
	"testing"
)

func getGameStateTest2() GameState {
	/*
		4 - - - -
		3 - 0 - f
		2 0 0 1 1
		1 x - 1 f
		0 - x - -
			0 1 2 3
	*/
	snakes := []Battlesnake{
		{
			ID:     "1",
			Health: 99,
			Head:   Coord{X: 0, Y: 2},
			Body:   []Coord{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 1, Y: 3}},
		},
		{
			ID:     "2",
			Health: 100,
			Head:   Coord{X: 3, Y: 2},
			Body:   []Coord{{X: 3, Y: 2}, {X: 2, Y: 2}, {X: 2, Y: 1}},
		},
	}

	food := []Coord{
		{X: 3, Y: 1},
		{X: 3, Y: 3},
	}

	hazards := []Coord{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	state := GameState{
		Turn: 2,
		You:  snakes[0],
		Board: Board{
			Snakes: []Battlesnake{
				snakes[0], snakes[1],
			},
			Food:    food,
			Hazards: hazards,
			Width:   4,
			Height:  4,
		},
	}

	return state
}

func randomPlayerTest(t *testing.T, state GameState, badMoves []string) {
	player := minimaxplayer.BuildRandomPlayer()
	move := player.Move(state)
	for _, badmove := range badMoves {
		if move == badmove {
			println("cannot go ", move)
			t.FailNow()
		}
	}
}

func Test_randomPlayer(t *testing.T) {
	randomPlayerTest(t, getGameStateTest1(), []string{"right", "up"})
	randomPlayerTest(t, getGameStateTest2(), []string{"left", "down", "right"})
}

func getRealTestData() []models.GameState {
	data := []string{
		"{\"game\":{\"id\":\"e68dab7c-e39d-4639-873a-5fa0466f4ec2\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\",\"settings\":{\"foodSpawnChance\":\n15,\"minimumFood\":1,\"hazardDamagePerTurn\":14}},\"map\":\"standard\",\"source\":\"\",\"timeout\":500},\"turn\":8,\"board\":{\"height\":11,\"width\":\n11,\"food\":[{\"x\":2,\"y\":0},{\"x\":5,\"y\":5}],\"hazards\":[],\"snakes\":[{\"id\":\"641ca4d3-db28-4bec-8123-b89be513c2fb\",\"name\":\"jacksnake1\",\n\"health\":100,\"body\":[{\"x\":0,\"y\":8},{\"x\":0,\"y\":9},{\"x\":0,\"y\":10},{\"x\":0,\"y\":10}],\"head\":{\"x\":0,\"y\":8},\"length\":4,\"latency\":\"0\",\"s\nhout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}},{\"id\":\"47da1b3b-7798-47b1-b6d5-b38470cc3d9b\",\"n\name\":\"jacksnake2\",\"health\":92,\"body\":[{\"x\":1,\"y\":7},{\"x\":1,\"y\":6},{\"x\":2,\"y\":6}],\"head\":{\"x\":1,\"y\":7},\"length\":3,\"latency\":\"1\",\"\nshout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}]},\"you\":{\"id\":\"641ca4d3-db28-4bec-8123-b89be51\n3c2fb\",\"name\":\"jacksnake1\",\"health\":100,\"body\":[{\"x\":0,\"y\":8},{\"x\":0,\"y\":9},{\"x\":0,\"y\":10},{\"x\":0,\"y\":10}],\"head\":{\"x\":0,\"y\":8},\n\"length\":4,\"latency\":\"0\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}}",
		"{\"game\":{\"id\":\"e68dab7c-e39d-4639-873a-5fa0466f4ec2\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\",\"settings\":{\"foodSpawnChance\":\n15,\"minimumFood\":1,\"hazardDamagePerTurn\":14}},\"map\":\"standard\",\"source\":\"\",\"timeout\":500},\"turn\":7,\"board\":{\"height\":11,\"width\":\n11,\"food\":[{\"x\":0,\"y\":8},{\"x\":2,\"y\":0},{\"x\":5,\"y\":5}],\"hazards\":[],\"snakes\":[{\"id\":\"641ca4d3-db28-4bec-8123-b89be513c2fb\",\"name\"\n:\"jacksnake1\",\"health\":93,\"body\":[{\"x\":0,\"y\":9},{\"x\":0,\"y\":10},{\"x\":1,\"y\":10}],\"head\":{\"x\":0,\"y\":9},\"length\":3,\"latency\":\"0\",\"sh\nout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}},{\"id\":\"47da1b3b-7798-47b1-b6d5-b38470cc3d9b\",\"na\nme\":\"jacksnake2\",\"health\":93,\"body\":[{\"x\":1,\"y\":6},{\"x\":2,\"y\":6},{\"x\":2,\"y\":5}],\"head\":{\"x\":1,\"y\":6},\"length\":3,\"latency\":\"0\",\"s\nhout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}]},\"you\":{\"id\":\"641ca4d3-db28-4bec-8123-b89be513\nc2fb\",\"name\":\"jacksnake1\",\"health\":93,\"body\":[{\"x\":0,\"y\":9},{\"x\":0,\"y\":10},{\"x\":1,\"y\":10}],\"head\":{\"x\":0,\"y\":9},\"length\":3,\"late\nncy\":\"0\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}}",
	}
	result := []models.GameState{}

	for _, str := range data {
		var state GameState
		json.Unmarshal([]byte(str), &state)
		result = append(result, state)
		fmt.Printf("%+v", state) // state is not getting set
	}

	return result
}

// integrationTest
func Test_PlayerCanRespondToMultipleRequests(t *testing.T) {
	testData := getRealTestData()
	player := minimaxplayer.BuildMinimaxPlayer()

	for _, state := range testData {
		player.Move(state)
	}
}
