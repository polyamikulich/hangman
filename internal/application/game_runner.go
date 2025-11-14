package application

import (
	"fmt"
	"os"
	"strings"
)

// RunInteractiveMode запускает интерактивный режим игры.
// Принимает интерфейсы для ввода, получения слов и отображения информации
func RunInteractiveMode(inputReader InputReader, wordRepo WordRepository, ui UI) error {
	fmt.Println("Добро пожаловать в игру «Виселица»!")

	// Выбираем категорию слова
	category := inputReader.ChooseCategory()
	difficulty := inputReader.ChooseWordDifficulty()
	maxMistakes := inputReader.ChooseGameDifficulty()
	isHintUsed := inputReader.ChooseHint()

	// Получаем случайное слово из выбранной категории
	word, err := wordRepo.GetRandomWord(category, difficulty)
	if err != nil {
		return fmt.Errorf("Не удалось выбрать слово: %w", err)
	}

	// Создаём новую игровую сессию
	session := NewGameSession(word, maxMistakes)

	// Основной игровой цикл
	for {
		status := session.GetStatus(word)

		if status.Finished {
			// Финальная визуализация
			ui.DrawHangman(session.ReturnCurrCountMistakes(), session.ReturnMaxCountMistakes())
			ui.DrawCurrentState(session.GetCurrentState())

			// Показываем результат
			if session.IsWin() {
				ui.ShowWin()
			} else {
				ui.ShowLoss(word.Value())
			}
			break
		}

		// Рисуем текущее состояние
		ui.DrawHangman(session.ReturnCurrCountMistakes(), session.ReturnMaxCountMistakes())
		ui.DrawCurrentState(session.GetCurrentState())
		ui.DrawAttempts(session.RemainingMistakes())

		if isHintUsed == 1 {
			ui.DrawHint(word.Hint())
		}

		// Ждём ввода буквы
		fmt.Println()
		fmt.Println("Введите букву")
		letter := inputReader.ReadRune()

		// Обрабатываем угадывание
		session.GuessingLetter(letter)

		fmt.Println() // пустая строка для читаемости
	}

	return nil
}

// RunTestMode запускает тестовый режим игры.
// Принимает два слова: загаданное и угаданное
func RunTestMode(hiddenWord, guessedWord string) error {
	// Проверка слов на длину
	if len(hiddenWord) != len(guessedWord) {
		return fmt.Errorf("Слова в arg1 и args2 должны состоять из одинакового количества знаков")
	}

	orig_hiddenRune := []rune(hiddenWord)

	hiddenWord = strings.ToLower(hiddenWord)
	guessedWord = strings.ToLower(guessedWord)

	hiddenRune := []rune(hiddenWord)
	guessedRune := []rune(guessedWord)

	resultRunes := make([]rune, len(hiddenRune))

	// Проверка каждой буквы
	for i := 0; i < len(hiddenRune); i++ {
		if hiddenRune[i] == guessedRune[i] {
			resultRunes[i] = orig_hiddenRune[i]
		} else {
			resultRunes[i] = '*'
		}
	}

	outcome := "POS"
	if strings.ContainsRune(string(resultRunes), '*') {
		outcome = "NEG"
	}

	// Показываем результат
	fmt.Printf("%s;%s\n", string(resultRunes), outcome)
	return nil
}

// ShowUsage отображает инструкцию по использованию программы
func ShowUsage() {
	fmt.Fprintf(os.Stderr, "Использование:\n")
	fmt.Fprintf(os.Stderr, "  %s — запуск в интерактивном режиме\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s <слово> <угадывание> — тестовый режим\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Пример:\n")
	fmt.Fprintf(os.Stderr, "  %s волокно толокно — выведет: *олокно;NEG\n", os.Args[0])
}
