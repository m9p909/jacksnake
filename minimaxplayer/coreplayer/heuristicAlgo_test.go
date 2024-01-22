package coreplayer_test

import (
	"jacksnake/minimaxplayer/coreplayer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsdfStuff(t *testing.T) {
	res := coreplayer.BuildProductOfDirections(4)
	assert.Len(t, res, 256)
}
