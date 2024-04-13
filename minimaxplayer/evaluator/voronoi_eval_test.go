package evaluator

import (
	"jacksnake/minimaxplayer/coreplayer"
	"reflect"
	"testing"
)

func TestVoronoiScore(t *testing.T) {
	type args struct {
		board *coreplayer.GameBoard
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "first",
			args: struct{ board *coreplayer.GameBoard }{
				board: &coreplayer.GameBoard{
					Height: 11,
					Width:  11,
					Turn:   0,
					Food:   make([]coreplayer.Point, 0),
					Snakes: []coreplayer.Snake{
						{
							ID:     0,
							Health: 100,
							Body: []coreplayer.Point{
								{
									X: 0,
									Y: 2,
								},
								{
									X: 1,
									Y: 2,
								},
								{
									X: 2,
									Y: 2,
								},
							},
						},
						{
							ID:     1,
							Health: 100,
							Body: []coreplayer.Point{
								{
									X: 10,
									Y: 10,
								},
								{
									X: 10,
									Y: 9,
								},
								{
									X: 9,
									Y: 9,
								},
							},
						},
					},
					Hazards: make([]coreplayer.Point, 0),
				}},
			want: []float64{0.487, 0.363, 0, 0},
		},
		{
			name: "first",
			args: struct{ board *coreplayer.GameBoard }{
				board: &coreplayer.GameBoard{
					Height: 11,
					Width:  11,
					Turn:   0,
					Food:   make([]coreplayer.Point, 0),
					Snakes: []coreplayer.Snake{
						{
							ID:     0,
							Health: 100,
							Body: []coreplayer.Point{
								{
									X: 10,
									Y: 10,
								},
								{
									X: 10,
									Y: 9,
								},
								{
									X: 9,
									Y: 9,
								},
							},
						},
						{
							ID:     1,
							Health: 100,
							Body: []coreplayer.Point{
								{
									X: 0,
									Y: 2,
								},
								{
									X: 1,
									Y: 2,
								},
								{
									X: 2,
									Y: 2,
								},
							},
						},
					},
					Hazards: make([]coreplayer.Point, 0),
				}},
			want: []float64{0.363, 0.487, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VoronoiScore(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VoronoiScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
