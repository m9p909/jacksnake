package minimaxplayer_test

import (
	"encoding/json"
	"fmt"
	"jacksnake/minimaxplayer"
	"jacksnake/models"
	. "jacksnake/models"
	"os"
	"testing"
	"time"

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

/*
func Test_randomPlayer(t *testing.T) {
	randomPlayerTest(t, getGameStateTest1(), []string{"right", "up"})
	randomPlayerTest(t, getGameStateTest2(), []string{"left", "down", "right"})
}
*/

type TestData struct {
	Data  models.GameState //`json:"data"`
	Name  string           //`json:"name"`
	Moves []string         //`json:"moves"`
}

type DataObject struct {
	Tests []TestData //`json:"tests"`
}

func getRealTestData() []TestData {
	result := DataObject{}
	raw, _ := os.ReadFile("testdata.json")
	err := json.Unmarshal(raw, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result.Tests
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

// control, used to determin if integration test is failing on async issue or normal issue
func Test_playerCanRespondToMultipleRequestsControlSync(t *testing.T) {
	testData := getRealTestData()
	player := minimaxplayer.BuildMinimaxPlayer()

	for i, data := range testData {

		timer := time.Now()
		res := player.Move(data.Data)
		since := time.Since(timer)
		val := assert.Less(t, since.Milliseconds(), int64(400))
		if !val {
			println(data.Name)
		}
		val = assert.Contains(t, testData[i].Moves, res)
		if !val {
			println(data.Name)
		}

	}
}

// integrationTest

func playerCanRespondToMultipleRequests(t *testing.T) {
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
		}(data.Data, chans[i])
	}

	result := make([]string, len(testData))

	for i := range chans {
		result[i] = <-chans[i]
	}

	for i := range result {
		assert.Contains(t, testData[i].Moves, result[i])
	}
}
