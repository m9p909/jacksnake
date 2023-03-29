package coreplayer_test

import (
	"encoding/json"
	. "jacksnake/minimaxplayer/coreplayer"
	"jacksnake/minimaxplayer/evaluator"
	"testing"
)

// integration
type TestData struct {
	board   GameBoard
	snakeId SnakeID
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
		{
			board:   "{\"Turn\":42,\"Height\":11,\"Width\":11,\"Food\":[{\"X\":2,\"Y\":0},{\"X\":5,\"Y\":5},{\"X\":4,\"Y\":10},{\"X\":1,\"Y\":2}],\"Snakes\":[{\"ID\":\"e52984ec-a077-4faf-b454-e8f7b26789ff\",\"Body\":[{\"X\":6,\"Y\":10},{\"X\":6,\"Y\":9},{\"X\":7,\"Y\":9},{\"X\":7,\"Y\":8},{\"X\":8,\"Y\":8}],\"Health\":70},{\"ID\":\"1ec3f9a8-4147-4b40-8c5c-79b6017127c2\",\"Body\":[{\"X\":5,\"Y\":9},{\"X\":5,\"Y\":8},{\"X\":4,\"Y\":8}],\"Health\":58}],\"Hazards\":[]}",
			snakeId: "e52984ec-a077-4faf-b454-e8f7b26789ff",
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
	algo := NewMinimaxAlgoMove(officialrulesapi.GetOfficialRules(), evaluator.NewSimpleEvaluator(), 5)
	for _, play := range data {
		algo.Move(play.board, play.snakeId)
	}
}
