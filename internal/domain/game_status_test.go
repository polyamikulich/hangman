package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameStatus_InProgressGame(t *testing.T) {
	status := GameStatus{
		Finished:          false,
		Win:               false,
		Loss:              false,
		Word:              "тест",
		CurrMistakes:      2,
		MaxMistakes:       5,
		State:             "т**т",
		RemainingMistakes: 3,
	}

	assert.False(t, status.Finished)
	assert.False(t, status.Win)
	assert.False(t, status.Loss)
	assert.Equal(t, "тест", status.Word)
	assert.Equal(t, 2, status.CurrMistakes)
	assert.Equal(t, 5, status.MaxMistakes)
	assert.Equal(t, "т**т", status.State)
	assert.Equal(t, 3, status.RemainingMistakes)
}

func TestGameStatus_DefaultValues(t *testing.T) {
	status := GameStatus{}

	assert.False(t, status.Finished)
	assert.False(t, status.Win)
	assert.False(t, status.Loss)
	assert.Equal(t, "", status.Word)
	assert.Equal(t, "", status.State)
	assert.Equal(t, 0, status.CurrMistakes)
	assert.Equal(t, 0, status.MaxMistakes)
	assert.Equal(t, 0, status.RemainingMistakes)
}

func TestGameStatus_WinGame(t *testing.T) {
	status := GameStatus{
		Finished:          true,
		Win:               true,
		Loss:              false,
		Word:              "победа",
		CurrMistakes:      2,
		MaxMistakes:       6,
		State:             "победа",
		RemainingMistakes: 4,
	}

	assert.True(t, status.Finished)
	assert.True(t, status.Win)
	assert.False(t, status.Loss)
	assert.Equal(t, "победа", status.Word)
	assert.Equal(t, 2, status.CurrMistakes)
	assert.Equal(t, 6, status.MaxMistakes)
	assert.Equal(t, "победа", status.State)
	assert.Equal(t, 4, status.RemainingMistakes)
}

func TestNewGameStatus_LossGame(t *testing.T) {
	status := GameStatus{
		Finished:          true,
		Win:               false,
		Loss:              true,
		Word:              "проигрыш",
		CurrMistakes:      6,
		MaxMistakes:       6,
		State:             "п*о***ры*",
		RemainingMistakes: 0,
	}

	assert.True(t, status.Finished)
	assert.False(t, status.Win)
	assert.True(t, status.Loss)
	assert.Equal(t, "проигрыш", status.Word)
	assert.Equal(t, 6, status.CurrMistakes)
	assert.Equal(t, 6, status.MaxMistakes)
	assert.Equal(t, "п*о***ры*", status.State)
	assert.Equal(t, 0, status.RemainingMistakes)
}
