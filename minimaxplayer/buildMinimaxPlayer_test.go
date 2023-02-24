package minimaxplayer_test

import (
	"encoding/json"
	"jacksnake/minimaxplayer"
	"jacksnake/models"
	. "jacksnake/models"
	"testing"

	"github.com/stretchr/testify/assert"
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

type TestData struct {
	data   models.GameState
	result []string
}

type RawTestData struct {
	data   string
	result []string
}

func getRealTestData() []TestData {
	rawDates := []RawTestData{{
		"{\"game\":{\"id\":\"aa5568d6-5c84-4564-8cbf-f69151334b81\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\",\"settings\":{\"foodSpawnChance\":15,\"minimumFood\":1,\"hazardDamagePerTurn\":14}},\"map\":\"standard\",\"source\":\"\",\"timeout\":500},\"turn\":43,\"board\":{\"height\":11,\"width\":11,\"food\":[{\"x\":6,\"y\":0},{\"x\":8,\"y\":6},{\"x\":8,\"y\":4},{\"x\":7,\"y\":5}],\"hazards\":[],\"snakes\":[{\"id\":\"d58b8a75-38af-4d9a-9b07-9a5eeab76d0c\",\"name\":\"jacksnake2\",\"health\":59,\"body\":[{\"x\":3,\"y\":8},{\"x\":3,\"y\":9},{\"x\":3,\"y\":10},{\"x\":4,\"y\":10}],\"head\":{\"x\":3,\"y\":8},\"length\":4,\"latency\":\"52\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}},{\"id\":\"cf9f556a-9fd9-4f0b-93ba-89e6bf0c4bc0\",\"name\":\"jacksnake1\",\"health\":66,\"body\":[{\"x\":1,\"y\":8},{\"x\":1,\"y\":7},{\"x\":1,\"y\":6},{\"x\":0,\"y\":6},{\"x\":0,\"y\":5}],\"head\":{\"x\":1,\"y\":8},\"length\":5,\"latency\":\"57\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}]},\"you\":{\"id\":\"d58b8a75-38af-4d9a-9b07-9a5eeab76d0c\",\"name\":\"jacksnake2\",\"health\":59,\"body\":[{\"x\":3,\"y\":8},{\"x\":3,\"y\":9},{\"x\":3,\"y\":10},{\"x\":4,\"y\":10}],\"head\":{\"x\":3,\"y\":8},\"length\":4,\"latency\":\"0\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}}\r\n",
		[]string{"down", "right"},
	}, {
		"{\"game\":{\"id\":\"aa5568d6-5c84-4564-8cbf-f69151334b81\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\",\"settings\":{\"foodSpawnChance\":15,\"minimumFood\":1,\"hazardDamagePerTurn\":14}},\"map\":\"standard\",\"source\":\"\",\"timeout\":500},\"turn\":43,\"board\":{\"height\":11,\"width\":11,\"food\":[{\"x\":6,\"y\":0},{\"x\":8,\"y\":6},{\"x\":8,\"y\":4},{\"x\":7,\"y\":5}],\"hazards\":[],\"snakes\":[{\"id\":\"d58b8a75-38af-4d9a-9b07-9a5eeab76d0c\",\"name\":\"jacksnake2\",\"health\":59,\"body\":[{\"x\":3,\"y\":8},{\"x\":3,\"y\":9},{\"x\":3,\"y\":10},{\"x\":4,\"y\":10}],\"head\":{\"x\":3,\"y\":8},\"length\":4,\"latency\":\"52\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}},{\"id\":\"cf9f556a-9fd9-4f0b-93ba-89e6bf0c4bc0\",\"name\":\"jacksnake1\",\"health\":66,\"body\":[{\"x\":1,\"y\":8},{\"x\":1,\"y\":7},{\"x\":1,\"y\":6},{\"x\":0,\"y\":6},{\"x\":0,\"y\":5}],\"head\":{\"x\":1,\"y\":8},\"length\":5,\"latency\":\"57\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}]},\"you\":{\"id\":\"cf9f556a-9fd9-4f0b-93ba-89e6bf0c4bc0\",\"name\":\"jacksnake1\",\"health\":66,\"body\":[{\"x\":1,\"y\":8},{\"x\":1,\"y\":7},{\"x\":1,\"y\":6},{\"x\":0,\"y\":6},{\"x\":0,\"y\":5}],\"head\":{\"x\":1,\"y\":8},\"length\":5,\"latency\":\"0\",\"shout\":\"\",\"customizations\":{\"color\":\"#b13859\",\"head\":\"default\",\"tail\":\"default\"}}}\r\n",
		[]string{"left", "up"},
	}}
	result := []TestData{}

	for _, rawDate := range rawDates {
		var state GameState
		json.Unmarshal([]byte(rawDate.data), &state)

		result = append(result, TestData{
			data:   state,
			result: rawDate.result,
		})
	}

	return result
}

func stringInArray(arr []string, str string) bool {
	status := false
	for _, value := range arr {
		if str == value {
			status = true
		}
	}
	return status
}

// integrationTest
func Test_playerCanRespondToMultipleRequests(t *testing.T) {
	testData := getRealTestData()
	player := minimaxplayer.BuildMinimaxPlayer()

	chans := make([]chan string, len(testData))
	for i := range chans {
		chans[i] = make(chan string)
	}
	for i, data := range testData {
		go func(state GameState, out chan string) {
			res := player.Move(state)
			out <- res
		}(data.data, chans[i])
	}

	result := make([]string, len(testData))

	for i := range chans {
		result[i] = <-chans[i]
	}

	for i := range result {
		assert.Contains(t, testData[i].result, result[i])
	}
}
