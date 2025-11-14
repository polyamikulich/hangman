package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// В этом файле мы просто проверяем, что функции не паникуют при
// различных данных

// TestDrawHangman проверяет отображение виселицы
func TestDrawHangman(t *testing.T) {
	tests := []struct {
		name         string
		currMistakes int
		maxMistakes  int
	}{
		{
			name:         "No mistakes",
			currMistakes: 0,
			maxMistakes:  6,
		},
		{
			name:         "Some mistakes",
			currMistakes: 3,
			maxMistakes:  5,
		},
		{
			name:         "Maximum mistakes",
			currMistakes: 6,
			maxMistakes:  6,
		},
		{
			name:         "Scaled mistakes",
			currMistakes: 2,
			maxMistakes:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				DrawHangman(tt.currMistakes, tt.maxMistakes)
			})
		})
	}
}

// TestDrawCurrentState проверяет отображение текущего состояния слова
func TestDrawCurrentState(t *testing.T) {
	tests := []struct {
		name  string
		state string
	}{
		{
			name:  "Empty state",
			state: "",
		},
		{
			name:  "Partial state",
			state: "*е**",
		},
		{
			name:  "Complete state",
			state: "тест",
		},
		{
			name:  "State full",
			state: "****",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				DrawCurrentState(tt.state)
			})
		})
	}
}

// TestDrawAttempts проверяет отображение количества оставшихся попыток
func TestDrawAttempts(t *testing.T) {
	tests := []struct {
		name     string
		mistakes int
	}{
		{
			name:     "No attempts left",
			mistakes: 0,
		},
		{
			name:     "Some attempts left",
			mistakes: 3,
		},
		{
			name:     "Many attempts left",
			mistakes: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				DrawAttempts(tt.mistakes)
			})
		})
	}
}

// TestShowWin проверяет отображение сообщения о победе
func TestShowWin(t *testing.T) {
	assert.NotPanics(t, func() {
		ShowWin()
	})
}

// TestShowLoss проверяет отображение сообщения о поражении
func TestShowLoss(t *testing.T) {
	tests := []struct {
		name string
		word string
	}{
		{
			name: "Simple word",
			word: "тест",
		},
		{
			name: "Complex word",
			word: "сложное слово",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				ShowLoss(tt.word)
			})
		})
	}
}
