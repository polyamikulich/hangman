package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain"
)

func TestNewGameSession(t *testing.T) {
	word := domain.NewWord("пример", "категория", 1, "подсказка")
	session := NewGameSession(word, 5)

	assert.Equal(t, 0, session.ReturnCurrCountMistakes())
	assert.Equal(t, 5, session.ReturnMaxCountMistakes())
	assert.Equal(t, 5, session.RemainingMistakes())
}

func TestGameSession_GuessingLetter_Correct(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	result := session.GuessingLetter('т')

	assert.True(t, result)
	assert.Equal(t, 0, session.ReturnCurrCountMistakes())
}

func TestGameSession_GuessingLetter_Incorrect(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	result := session.GuessingLetter('а')

	assert.False(t, result)
	assert.Equal(t, 1, session.ReturnCurrCountMistakes())
}

func TestGameSession_GuessingLetter_DuplicateCorrect(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Угадываем букву первый раз
	firstTry := session.GuessingLetter('т')

	// Угадываем ту же букву второй раз
	secondTry := session.GuessingLetter('т')

	assert.True(t, firstTry)                              // Первая попытка успешна
	assert.False(t, secondTry)                            // Вторая попытка неуспешна (дубликат)
	assert.Equal(t, 0, session.ReturnCurrCountMistakes()) // Не увеличивается при повторе
}

func TestGameSession_GuessingLetter_DuplicateInorrect(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Угадываем букву первый раз
	firstTry := session.GuessingLetter('к') // Первая попытка неуспешна
	assert.False(t, firstTry)
	assert.Equal(t, 1, session.ReturnCurrCountMistakes())

	// Угадываем ту же букву второй раз
	secondTry := session.GuessingLetter('к')
	assert.False(t, secondTry)                            // Вторая попытка неуспешна (дубликат)
	assert.Equal(t, 1, session.ReturnCurrCountMistakes()) // Не увеличивается при повторе
}

func TestGameSession_GuessingLetter_CaseInsensitive(t *testing.T) {
	word := domain.NewWord("Тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Проверяем, что регистр не важен
	result := session.GuessingLetter('т') // строчная
	assert.True(t, result)
	assert.Equal(t, 0, session.ReturnCurrCountMistakes())

	result = session.GuessingLetter('Е') // заглавная
	assert.True(t, result)
	assert.Equal(t, 0, session.ReturnCurrCountMistakes())
}

func TestGameSession_GetCurrentState_InitialState(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Проверяем начальное состояние - все буквы скрыты
	state := session.GetCurrentState()
	assert.Equal(t, "****", state)
}

func TestGameSession_GetCurrentState_PartialGuess(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Угадываем одну букву
	session.GuessingLetter('е')
	state := session.GetCurrentState()
	assert.Equal(t, "*е**", state)
}

func TestGameSession_GetCurrentState_WordWithRepeatedLetters(t *testing.T) {
	word := domain.NewWord("балалайка", "музыка", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Угадываем букву 'а', которая встречается несколько раз
	session.GuessingLetter('а')
	state := session.GetCurrentState()
	assert.Equal(t, "*а*а*а**а", state) // Все 'а' открыты
}

func TestGameSession_GetCurrentState_AllGuessed(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Угадываем все уникальные буквы
	session.GuessingLetter('т')
	session.GuessingLetter('е')
	session.GuessingLetter('с')

	state := session.GetCurrentState()
	assert.Equal(t, "тест", state) // Все буквы открыты
}

func TestGameSession_GetCurrentState_NoLettersGuessed(t *testing.T) {
	word := domain.NewWord("программа", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	session.GuessingLetter('с')
	session.GuessingLetter('ф')

	// Не угадываем ни одной буквы
	state := session.GetCurrentState()
	assert.Equal(t, "*********", state) // Все буквы скрыты
}

func TestGameSession_GetCurrentState_CaseInsensitiveWordDisplay(t *testing.T) {
	word := domain.NewWord("ТеСт", "пример", 1, "подсказка") // Смешанный регистр в слове
	session := NewGameSession(word, 5)

	// Угадываем буквы в нижнем регистре
	session.GuessingLetter('т')
	session.GuessingLetter('с')

	state := session.GetCurrentState()
	assert.Equal(t, "Т*Ст", state) // Отображается в оригинальном регистре
}

func TestGameSession_IsWin(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Пока не угадали все буквы
	assert.False(t, session.IsWin())

	// Угадываем все буквы по одной
	session.GuessingLetter('т')
	assert.False(t, session.IsWin()) // Еще не все буквы

	session.GuessingLetter('е')
	assert.False(t, session.IsWin()) // Еще не все буквы

	session.GuessingLetter('с')
	assert.True(t, session.IsWin()) // Угадали все уникальные буквы
}

func TestGameSession_IsLoss(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 1) // Только одна попытка

	// Пока не превышено количество ошибок
	assert.False(t, session.IsLoss())

	// Делаем ошибку
	session.GuessingLetter('а')

	assert.True(t, session.IsLoss())
}

func TestGameSession_RemainingMistakes(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	assert.Equal(t, 5, session.RemainingMistakes())

	session.GuessingLetter('а') // Ошибка
	assert.Equal(t, 4, session.RemainingMistakes())

	session.GuessingLetter('б') // Ошибка
	assert.Equal(t, 3, session.RemainingMistakes())

	session.GuessingLetter('т')                     // Угадали
	assert.Equal(t, 3, session.RemainingMistakes()) // Ошибки не прибавляются
}

func TestGameSession_ReturnCountMistakes(t *testing.T) {
	word := domain.NewWord("тест", "пример", 1, "подсказка")
	session := NewGameSession(word, 5)

	// Проверяем начальные значения
	assert.Equal(t, 0, session.ReturnCurrCountMistakes())
	assert.Equal(t, 5, session.ReturnMaxCountMistakes())

	// Делаем ошибку
	session.GuessingLetter('а')

	// Проверяем после ошибки
	assert.Equal(t, 1, session.ReturnCurrCountMistakes())
	assert.Equal(t, 5, session.ReturnMaxCountMistakes())
}
