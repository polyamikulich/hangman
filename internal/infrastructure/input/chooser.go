package input

import (
	"fmt"
	"math/rand"
	"os"
)

func ChooseCategory() string {
	fmt.Println("\n Выберите категорию слова (введите цифру):")
	fmt.Println("1) Животные")
	fmt.Println("2) Фрукты")
	fmt.Println("3) Профессии")
	fmt.Println("4) Страны")
	fmt.Println("5) Музыкальные инструменты")
	fmt.Println("6) Любая случайная")

	num, err := ReadNumber(6)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка ввода: %v\n", err)
		fmt.Fprintln(os.Stderr, "Не удалось прочитать выбор. Завершение.")
		os.Exit(1)
	}

	switch num {
	case 1:
		return "животные"
	case 2:
		return "фрукты"
	case 3:
		return "профессии"
	case 4:
		return "страны"
	case 5:
		return "музыкальные инструменты"
	case 6:
		// Случайный выбор из доступных категорий
		categories := []string{"животные", "фрукты", "профессии", "страны", "музыкальные инструменты"}
		return categories[rand.Intn(len(categories))]
	default:
		fmt.Fprintln(os.Stderr, "Неверный номер. Выберите число от 1 до 6.")
		os.Exit(1)
		return ""
	}
}

func ChooseWordDifficulty() int {
	fmt.Println("\n Выберите сложность слова (введите цифру):")
	fmt.Println("1) Низкая")
	fmt.Println("2) Средняя")
	fmt.Println("3) Высокая")
	fmt.Println("4) Любая случайная")

	num, err := ReadNumber(4)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка ввода: %v\n", err)
		fmt.Fprintln(os.Stderr, "Не удалось прочитать выбор. Завершение.")
		os.Exit(1)
	}

	switch num {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		difficulty := []int{1, 2, 3}
		return difficulty[rand.Intn(len(difficulty))]
	default:
		fmt.Fprintln(os.Stderr, "Неверный номер. Выберите число от 1 до 4.")
		os.Exit(1)
		return -1
	}
}

func ChooseGameDifficulty() int {
	fmt.Println("\n Выберите сложность игры - количество допустимых ошибок (введите цифру):")
	fmt.Println("1) Очень сложно")
	fmt.Println("2) Сложно")
	fmt.Println("3) Достаточно сложно")
	fmt.Println("4) Средне")
	fmt.Println("5) Достаточно легко")
	fmt.Println("6) Легко (классическая версия игры)")
	fmt.Println("7) Любая случайная")

	num, err := ReadNumber(7)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка ввода: %v\n", err)
		fmt.Fprintln(os.Stderr, "Не удалось прочитать выбор. Завершение.")
		os.Exit(1)
	}

	switch num {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return 4
	case 5:
		return 5
	case 6:
		return 6
	case 7:
		difficulty := []int{1, 2, 3}
		return difficulty[rand.Intn(len(difficulty))]
	default:
		fmt.Fprintln(os.Stderr, "Неверный номер. Выберите число от 1 до 7.")
		os.Exit(1)
		return -1
	}
}

func ChooseHint() int {
	fmt.Println("\n Решите, нужна ли вам подсказка (введите цифру):")
	fmt.Println("1) Да, нужна")
	fmt.Println("2) Нет, не нужна")

	num, err := ReadNumber(2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка ввода: %v\n", err)
		fmt.Fprintln(os.Stderr, "Не удалось прочитать выбор. Завершение.")
		os.Exit(1)
	}

	switch num {
	case 1:
		return 1
	case 2:
		return 2
	default:
		fmt.Fprintln(os.Stderr, "Неверный номер. Выберите число от 1 до 2.")
		os.Exit(1)
		return -1
	}
}
