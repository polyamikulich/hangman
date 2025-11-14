package ui

import "fmt"

// Различные состояния виселицы
var hangmanStages = []string{
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

// DrawHangman выводит текущее состояние виселицы, масштабируя визуализацию
// под любое количество максимально допустимых ошибок (максимум 6)
func DrawHangman(currMistakes, maxMistakes int) {
	const totalStages = 6

	stageIndex := (totalStages * currMistakes) / maxMistakes

	fmt.Println(hangmanStages[stageIndex])
}

// DrawCurrentState выводит текущее состояние угадываемого слова
func DrawCurrentState(state string) {
	fmt.Printf("Слово: %s\n", state)
}

func DrawHint(hint string) {
	fmt.Printf("Подсказка: %s\n", hint)
}

// DrawAttempts выводит количество оставшихся ошибок
func DrawAttempts(mistakes int) {
	fmt.Printf("Осталось попыток: %d\n", mistakes)
}

// ShowWin выводит сообщение о победе
func ShowWin() {
	fmt.Println("Поздравляем! Вы угадали слово!")
}

// ShowLoss выводит сообщение о поражении
func ShowLoss(word string) {
	fmt.Printf("Вы проиграли. Загаданное слово: %s\n", word)
}
