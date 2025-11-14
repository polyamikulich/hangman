package application

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain"
)

// InputReader определяет интерфейс для чтения ввода пользователя
type InputReader interface {
	ChooseCategory() string

	ChooseGameDifficulty() int

	ChooseWordDifficulty() int

	ChooseHint() int

	ReadRune() rune

	ReadNumber(maxnum int) (int, error)
}

// WordRepository определяет интерфейс для получения случайного слова
type WordRepository interface {
	GetRandomWord(category string, difficulty int) (domain.Word, error)
}

// UI определяет интерфейс для вывода данных
type UI interface {
	DrawHangman(currMistakes, maxMistakes int)

	DrawCurrentState(state string)

	DrawAttempts(mistakes int)

	DrawHint(hint string)

	ShowWin()

	ShowLoss(word string)
}
