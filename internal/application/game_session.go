package application

import (
	"strings"
	"unicode"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain"
)

// GameSession представляет игровую сессию "Виселица"
type GameSession struct {
	hiddenWord        domain.Word
	guessedLetters    map[rune]bool
	maxCountMistakes  int
	currCountMistakes int
}

// NewGameSession создает новую игровую сессию
func NewGameSession(word domain.Word, maxCountMistakes int) *GameSession {
	return &GameSession{
		hiddenWord:        word,
		guessedLetters:    make(map[rune]bool),
		maxCountMistakes:  maxCountMistakes,
		currCountMistakes: 0,
	}
}

// GuessingLetter проверяет содержится ли буква в загаданном слове
// Буква добавляется в список угаданных.
// Возвращается булевское значение
func (gs *GameSession) GuessingLetter(letter rune) bool {
	letter = unicode.ToLower(letter)

	if gs.guessedLetters[letter] {
		return false
	}

	gs.guessedLetters[letter] = true

	// Проверка, есть ли буква в слове
	letterInWord := false
	for _, s := range gs.hiddenWord.Value() {
		if unicode.ToLower(s) == letter {
			letterInWord = true
			break
		}
	}

	if !letterInWord {
		gs.currCountMistakes += 1
	}

	return letterInWord
}

// GetCurrentState возвращает текущее состояние угадывания слова
// Угаданную букву выводим, неугаданную заменяем на *
func (gs *GameSession) GetCurrentState() string {
	var result []rune

	for _, letter := range gs.hiddenWord.Value() {
		lower_letter := unicode.ToLower(letter)
		if gs.guessedLetters[lower_letter] {
			result = append(result, letter)
		} else {
			result = append(result, '*')
		}
	}

	return string(result)
}

// IsWin определяет, выиграл ли игрок
// Победа = всё слово полностью угадано
func (gs *GameSession) IsWin() bool {
	currentState := gs.GetCurrentState()
	answer := strings.ContainsRune(currentState, '*')

	return !answer
}

// IsLoss определяет, проиграл ли игрок
// Проигрыш = ошибки превысили максимально допустимое число раз
func (gs *GameSession) IsLoss() bool {
	return gs.currCountMistakes >= gs.maxCountMistakes
}

// RemainingMistakes возвращает количество оставшихся ошибок
func (gs *GameSession) RemainingMistakes() int {
	return gs.maxCountMistakes - gs.currCountMistakes
}

// ReturnMaxCountMistakes возвращает максимально допустимое число ошибок
func (gs *GameSession) ReturnMaxCountMistakes() int {
	return gs.maxCountMistakes
}

// ReturnCurrCountMistakes возвращает текущее число ошибок
func (gs *GameSession) ReturnCurrCountMistakes() int {
	return gs.currCountMistakes
}

// GetStatus возвращает текущее состояние игры
func (gs *GameSession) GetStatus(word domain.Word) domain.GameStatus {
	return domain.GameStatus{
		Word:              word.Value(),
		State:             gs.GetCurrentState(),
		Win:               gs.IsWin(),
		Loss:              gs.IsLoss(),
		MaxMistakes:       gs.maxCountMistakes,
		CurrMistakes:      gs.currCountMistakes,
		Finished:          gs.IsWin() || gs.IsLoss(),
		RemainingMistakes: gs.RemainingMistakes(),
	}
}
