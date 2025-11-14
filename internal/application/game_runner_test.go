package application

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain"
)

// Моки для интерфейсов
type MockInputReader struct {
	mock.Mock
}

func (m *MockInputReader) ReadRune() rune {
	args := m.Called()
	return args.Get(0).(rune)
}

func (m *MockInputReader) ReadNumber(max int) (int, error) {
	args := m.Called(max)
	return args.Int(0), args.Error(1)
}

func (m *MockInputReader) ChooseCategory() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockInputReader) ChooseWordDifficulty() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockInputReader) ChooseGameDifficulty() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockInputReader) ChooseHint() int {
	args := m.Called()
	return args.Int(0)
}

type MockWordRepository struct {
	mock.Mock
}

func (m *MockWordRepository) GetRandomWord(category string, difficulty int) (domain.Word, error) {
	args := m.Called(category, difficulty)
	return args.Get(0).(domain.Word), args.Error(1)
}

type MockUI struct {
	mock.Mock
}

func (m *MockUI) DrawHangman(currMistakes, maxMistakes int) {
	m.Called(currMistakes, maxMistakes)
}

func (m *MockUI) DrawHint(hint string) {
	m.Called(hint)
}

func (m *MockUI) DrawCurrentState(state string) {
	m.Called(state)
}

func (m *MockUI) DrawAttempts(remainingMistakes int) {
	m.Called(remainingMistakes)
}

func (m *MockUI) ShowWin() {
	m.Called()
}

func (m *MockUI) ShowLoss(word string) {
	m.Called(word)
}

func TestRunInteractiveMode_Success_WithHint(t *testing.T) {
	// Создаем моки
	mockInput := new(MockInputReader)
	mockRepo := new(MockWordRepository)
	mockUI := new(MockUI)

	// Настраиваем ожидания для input
	mockInput.On("ChooseCategory").Return("животные")
	mockInput.On("ChooseWordDifficulty").Return(1)
	mockInput.On("ChooseGameDifficulty").Return(5)
	mockInput.On("ChooseHint").Return(1)
	mockInput.On("ReadRune").Return('т').Once()
	mockInput.On("ReadRune").Return('е').Once()
	mockInput.On("ReadRune").Return('с').Once()

	// Настраиваем ожидания для repo
	word := domain.NewWord("тест", "животные", 1, "пример подсказки")
	mockRepo.On("GetRandomWord", "животные", 1).Return(word, nil)

	// Настраиваем ожидания для UI
	mockUI.On("DrawHangman", mock.Anything, mock.Anything).Return()
	mockUI.On("DrawCurrentState", mock.Anything).Return()
	mockUI.On("DrawAttempts", mock.Anything).Return()
	mockUI.On("DrawHint", "пример подсказки").Return()
	mockUI.On("ShowWin").Return()

	// Выполняем тест
	err := RunInteractiveMode(mockInput, mockRepo, mockUI)

	// Проверяем результат
	assert.NoError(t, err)

	// Проверяем, что все методы были вызваны
	mockInput.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockUI.AssertExpectations(t)

	mockUI.On("DrawHint", "пример подсказки")
}

func TestRunInteractiveMode_Success_WithoutHint(t *testing.T) {
	// Создаем моки
	mockInput := new(MockInputReader)
	mockRepo := new(MockWordRepository)
	mockUI := new(MockUI)

	// Настраиваем ожидания для input
	mockInput.On("ChooseCategory").Return("животные")
	mockInput.On("ChooseWordDifficulty").Return(1)
	mockInput.On("ChooseGameDifficulty").Return(5)
	mockInput.On("ChooseHint").Return(0)
	mockInput.On("ReadRune").Return('т').Once()
	mockInput.On("ReadRune").Return('е').Once()
	mockInput.On("ReadRune").Return('с').Once()

	// Настраиваем ожидания для repo
	word := domain.NewWord("тест", "животные", 1, "пример подсказки")
	mockRepo.On("GetRandomWord", "животные", 1).Return(word, nil)

	// Настраиваем ожидания для UI
	mockUI.On("DrawHangman", mock.Anything, mock.Anything).Return()
	mockUI.On("DrawCurrentState", mock.Anything).Return()
	mockUI.On("DrawAttempts", mock.Anything).Return()
	mockUI.On("ShowWin").Return()

	// Выполняем тест
	err := RunInteractiveMode(mockInput, mockRepo, mockUI)

	// Проверяем результат
	assert.NoError(t, err)

	// Проверяем, что все методы были вызваны
	mockInput.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockUI.AssertExpectations(t)

	mockUI.AssertNotCalled(t, "DrawHint")
}

func TestRunInteractiveMode_Loss(t *testing.T) {
	// Создаем моки
	mockInput := new(MockInputReader)
	mockRepo := new(MockWordRepository)
	mockUI := new(MockUI)

	// Настраиваем ожидания для input
	mockInput.On("ChooseCategory").Return("животные")
	mockInput.On("ChooseWordDifficulty").Return(1)
	mockInput.On("ChooseGameDifficulty").Return(1) // Только одна попытка
	mockInput.On("ChooseHint").Return(1)
	mockInput.On("ReadRune").Return('а').Once() // Неправильная буква

	// Настраиваем ожидания для repo
	word := domain.NewWord("тест", "животные", 1, "подсказка")
	mockRepo.On("GetRandomWord", "животные", 1).Return(word, nil)

	// Настраиваем ожидания для UI
	mockUI.On("DrawHangman", 0, 1).Return()
	mockUI.On("DrawCurrentState", "****").Return()
	mockUI.On("DrawAttempts", 1).Return()
	mockUI.On("DrawHangman", 1, 1).Return()
	mockUI.On("DrawCurrentState", "****").Return()
	mockUI.On("DrawHint", "подсказка").Return()
	mockUI.On("ShowLoss", "тест").Return()

	// Выполняем тест
	err := RunInteractiveMode(mockInput, mockRepo, mockUI)

	// Проверяем результат
	assert.NoError(t, err)

	// Проверяем, что все методы были вызваны
	mockInput.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
	mockUI.AssertExpectations(t)
}

func TestRunInteractiveMode_WordRepositoryError(t *testing.T) {
	// Создаем моки
	mockInput := new(MockInputReader)
	mockRepo := new(MockWordRepository)
	mockUI := new(MockUI)

	// Настраиваем ожидания для input
	mockInput.On("ChooseCategory").Return("животные")
	mockInput.On("ChooseWordDifficulty").Return(1)
	mockInput.On("ChooseGameDifficulty").Return(5)
	mockInput.On("ChooseHint").Return(0)

	// Настраиваем ожидания для repo с ошибкой
	mockRepo.On("GetRandomWord", "животные", 1).Return(domain.Word{}, assert.AnError)

	// Выполняем тест
	err := RunInteractiveMode(mockInput, mockRepo, mockUI)

	// Проверяем результат
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Не удалось выбрать слово")

	// Проверяем, что UI не вызывался
	mockUI.AssertNotCalled(t, "DrawHangman")
	mockUI.AssertNotCalled(t, "DrawCurrentState")
	mockUI.AssertNotCalled(t, "DrawAttempts")
	mockUI.AssertNotCalled(t, "DrawHint")
	mockUI.AssertNotCalled(t, "ShowWin")
	mockUI.AssertNotCalled(t, "ShowLoss")
}

func TestRunTestMode_Error_DifferentLengths(t *testing.T) {
	// Тест ошибки при разных длинах слов
	err := RunTestMode("слово", "длинноеслово")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Слова в arg1 и args2 должны состоять из одинакового количества знаков")

	err = RunTestMode("длинноеслово", "слово")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Слова в arg1 и args2 должны состоять из одинакового количества знаков")
}

func TestRunTestMode_EmptyWords(t *testing.T) {
	// Оба слова пустые
	err := RunTestMode("", "")
	assert.NoError(t, err)

	// Одно слово пустое, другое нет
	err = RunTestMode("", "тест")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Слова в arg1 и args2 должны состоять из одинакового количества знаков")

	err = RunTestMode("тест", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Слова в arg1 и args2 должны состоять из одинакового количества знаков")
}

// TestRunTestMode_OutputFormat проверяет формат вывода в тестовом режиме
func TestRunTestMode_OutputFormat(t *testing.T) {
	tests := []struct {
		name           string
		hiddenWord     string
		guessedWord    string
		expectedOutput string
	}{
		{
			name:           "Complete match",
			hiddenWord:     "окно",
			guessedWord:    "окно",
			expectedOutput: "окно;POS\n",
		},
		{
			name:           "Partial match",
			hiddenWord:     "волокно",
			guessedWord:    "толокно",
			expectedOutput: "*олокно;NEG\n",
		},
		{
			name:           "No match",
			hiddenWord:     "волокно",
			guessedWord:    "барахло",
			expectedOutput: "******о;NEG\n",
		},
		{
			name:           "Case insensitive",
			hiddenWord:     "Тест",
			guessedWord:    "тест",
			expectedOutput: "Тест;POS\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Перенаправляем stdout для захвата вывода
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := RunTestMode(tt.hiddenWord, tt.guessedWord)
			assert.NoError(t, err)

			// Восстанавливаем stdout
			w.Close()
			os.Stdout = old

			// Читаем захваченный вывод
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}

func TestRunTestMode_SingleCharacterWords(t *testing.T) {
	// Совпадающие символы
	err := RunTestMode("а", "а")
	assert.NoError(t, err)

	// Не совпадающие символы
	err = RunTestMode("а", "б")
	assert.NoError(t, err)
}

func TestShowUsage(t *testing.T) {
	// ShowUsage просто выводит текст, проверим, что не паникует
	assert.NotPanics(t, func() {
		ShowUsage()
	})
}
