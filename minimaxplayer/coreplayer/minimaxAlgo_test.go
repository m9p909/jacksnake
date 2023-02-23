package coreplayer_test

import (
	"encoding/json"
	"fmt"
	. "jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/evaluator"
	"jacksnake/minimaxplayer/officialrulesapi"
	"testing"
)

// integration
type TestData struct {
	board   GameBoard
	snakeId string
}

type TestDataRaw struct {
	board   string
	snakeId string
}

func getRawTestData() []TestDataRaw {
	return []TestDataRaw{
		{
			board:   "{\"Turn\":17,\"Height\":11,\"Width\":11,\"Food\":[{\"X\":10,\"Y\":2},{\"X\":0,\"Y\":2},{\"X\":5,\"Y\":5},{\"X\":9,\"Y\":0}],\"Snakes\":[{\"ID\":\"e80f0120-711c-4770-a0d0-0dc792d32055\",\"Body\":[{\"X\":5,\"Y\":8},{\"X\":5,\"Y\":7},{\"X\":4,\"Y\":7}],\"Health\":83},{\"ID\":\"d6b5a4a9-4909-4bef-a236-ef41cc975a46\",\"Body\":[{\"X\":0,\"Y\":1},{\"X\":1,\"Y\":1},{\"X\":2,\"Y\":1},{\"X\":2,\"Y\":0}],\"Health\":93}],\"Hazards\":[]}\n",
			snakeId: "d6b5a4a9-4909-4bef-a236-ef41cc975a46",
		},
		{
			board:   "{\"Turn\":18,\"Height\":11,\"Width\":11,\"Food\":[{\"X\":10,\"Y\":2},{\"X\":5,\"Y\":5},{\"X\":9,\"Y\":0}],\"Snakes\":[{\"ID\":\"e80f0120-711c-4770-a0d0-0dc792d32055\",\"Body\":[{\"X\":5,\"Y\":9},{\"X\":5,\"Y\":8},{\"X\":5,\"Y\":7}],\"Health\":82},{\"ID\":\"d6b5a4a9-4909-4bef-a236-ef41cc975a46\",\"Body\":[{\"X\":0,\"Y\":2},{\"X\":0,\"Y\":1},{\"X\":1,\"Y\":1},{\"X\":2,\"Y\":1},{\"X\":2,\"Y\":1}],\"Health\":100}],\"Hazards\":[]}\n",
			snakeId: "d6b5a4a9-4909-4bef-a236-ef41cc975a46",
		},
	}
}

func rawTestDataToTestData(dataRaw TestDataRaw) TestData {
	var v1 GameBoard
	data := []byte(dataRaw.board)
	json.Unmarshal(data, &v1)
	return TestData{
		v1,
		dataRaw.snakeId,
	}
}

func getTestData() []TestData {
	raw := getRawTestData()
	res := []TestData{}
	for _, rawData := range raw {
		res = append(res, rawTestDataToTestData(rawData))
	}
	return res
}

func Test_MinimaxPlayer(t *testing.T) {
	data := getTestData()
	fmt.Printf("%+v", data)
	algo := NewMinimaxAlgoMove(officialrulesapi.GetOfficialRules(), evaluator.NewSimpleEvaluator(), 1)
	for _, play := range data {
		algo.Move(play.board, play.snakeId)
	}
}
