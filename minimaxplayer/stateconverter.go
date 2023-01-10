package minimaxplayer

import (
	"jacksnake/minimaxplayer/coreplayer"
	"jacksnake/models"
)

type StateConverterImpl struct{}

func (StateConverterImpl) StateToCore(state models.GameState) coreplayer.GameBoard {
	return coreplayer.GameBoard{}
}

func (StateConverterImpl) CoreToState(state coreplayer.GameBoard) models.GameState {
	return models.GameState{}
}
