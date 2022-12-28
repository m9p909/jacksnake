package evaluateboard

import (
	"fmt"
	. "jacksnake/models"
	"strconv"
	"testing"
)

func Test_getSnakes(t *testing.T) {
	snakes := []Battlesnake{{ID: "a"}, {ID: "b"}, {ID: "c"}}
	data := GameState{
		You: snakes[0],
		Board: Board{
			Snakes: []Battlesnake{
				snakes[1], snakes[2],
			},
		},
	}

	result := getSnakes(data)

	if result[0].ID != snakes[1].ID {
		t.Fail()
	}

	if result[1].ID != snakes[2].ID {
		t.Fail()
	}

	if result[2].ID != snakes[0].ID {
		t.Fail()
	}

}

func Test_buildSnakeBoard(t *testing.T) {

	snakes := []Battlesnake{
		{Body: []Coord{{X: 1, Y: 2}}},
		{Body: []Coord{{X: 2, Y: 2},
			{X: 2, Y: 1}}}}

	board := buildSnakeBoard(snakes, 5, 5)

	if board[2][1] != strconv.Itoa(0) {
		t.Fail()
		println(board[2][1])
		println("should be 0")

	}

	if board[2][2] != strconv.Itoa(1) || board[1][2] != strconv.Itoa(1) {
		t.Fail()
		println(board[2][2])
		println("should be 1")
		println(board[2][1])
		println("should be 1")
	}

	//displayBoard(board)

}

func Test_createDistanceGraph(t *testing.T) {
	snakes := []Battlesnake{
		{Head: Coord{X: 1, Y: 2},
			Body: []Coord{{X: 1, Y: 2}}},
		{Head: Coord{X: 2, Y: 2},
			Body: []Coord{{X: 2, Y: 2},
				{X: 2, Y: 1}}}}

	g := createDistanceGraph(buildSnakeBoard(snakes, 5, 5), snakes[0])
	expected := [][]int{
		{3, 2, 3, 4, 5},
		{2, 1, -1, 5, 6},
		{1, 0, -1, 4, 5},
		{2, 1, 2, 3, 4},
		{3, 2, 3, 4, 5},
	}

	for i := range expected {
		for j := range expected[i] {
			if !(expected[i][j] == g[i][j]) {
				t.Fail()
				fmt.Printf("i: %d, j: %d, expected %d, actual %d", i, j, expected[i][j], g[i][j])
			}
		}
	}

	/*
		for _, b := range g {
			for _, c := range b {
				if c == 2000 {
					c = -1
				}
				print("|")
				print(c)
				print("|")
			}
			print("\n")
		}
	*/
}

func Test_getFoodScore(t *testing.T) {
	snakes := []Battlesnake{
		{Head: Coord{X: 1, Y: 2},
			Body: []Coord{{X: 1, Y: 2}}},
		{Head: Coord{X: 2, Y: 2},
			Body: []Coord{{X: 2, Y: 2},
				{X: 2, Y: 1}}}}
	food := []Coord{{X: 2, Y: 3}, {X: 4, Y: 3}}

	v := getFoodScore(food, createDistanceGraph(buildSnakeBoard(snakes, 5, 5), snakes[0]))
	print(v)

}
