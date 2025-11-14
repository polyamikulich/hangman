package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// ReadRune читает у пользователя одну букву
func ReadRune() rune {
	return ReadRuneFromReader(os.Stdin)
}

// ReadRuneFromReader читает одну букву из заданного io.Reader
func ReadRuneFromReader(reader io.Reader) rune {
	scanner := bufio.NewScanner(reader)
	for {
		if !scanner.Scan() {
			// Ошибка ввода (например, Ctrl+D)
			fmt.Println()
			fmt.Println("Ошибка ввода. Завершение.")
			os.Exit(1)
		}

		text := scanner.Text()
		r, err := ProcessInputString(text)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		return unicode.ToLower(r)
	}
}

// ReadNumber читает у пользователя цифру
// в диапазоне от 1 до maxnum.
func ReadNumber(maxnum int) (int, error) {
	return ReadNumberFromReader(os.Stdin, maxnum)
}

// ReadNumberFromReader читает цифру из заданного io.Reader
func ReadNumberFromReader(reader io.Reader, maxnum int) (int, error) {
	scanner := bufio.NewScanner(reader)
	for {
		if !scanner.Scan() {
			// Ошибка ввода (например, Ctrl+D)
			fmt.Println()
			fmt.Println("Ошибка ввода. Завершение.")
			os.Exit(1)
		}

		// Удаляем пробелы по краям
		text := strings.TrimSpace(scanner.Text())
		num, err := ProcessNumberString(text, maxnum)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		return num, nil
	}
}

// ProcessInputString обрабатывает строку ввода для получения буквы
func ProcessInputString(text string) (rune, error) {
	trimmed := strings.TrimSpace(text)
	if len(trimmed) == 0 {
		return 0, &InputError{msg: "Ввод пустой. Попробуйте снова."}
	}

	runes := []rune(trimmed)
	if len(runes) != 1 {
		return 0, &InputError{msg: "Введите ровно одну букву."}
	}

	r := runes[0]
	if !unicode.IsLetter(r) {
		return 0, &InputError{msg: "Это не буква. Попробуйте снова."}
	}

	return r, nil
}

// ProcessNumberString обрабатывает строку ввода для получения числа
func ProcessNumberString(text string, maxnum int) (int, error) {
	trimmed := strings.TrimSpace(text)
	if len(trimmed) == 0 {
		return 0, &InputError{msg: "Ввод пустой. Попробуйте снова."}
	}

	runes := []rune(trimmed)
	if len(runes) != 1 {
		return 0, &InputError{msg: "Введите ровно одну цифру"}
	}

	r := runes[0]
	if !unicode.IsDigit(r) {
		return 0, &InputError{msg: "Это не цифра. Попробуйте снова."}
	}

	num := int(r - '0') // преобразуем '1' -> 1
	if num < 1 || num > maxnum {
		return 0, &InputError{msg: fmt.Sprintf("Номер должен быть от 1 до %d", maxnum)}
	}

	return num, nil
}

// InputError представляет ошибку ввода
type InputError struct {
	msg string
}

func (e *InputError) Error() string {
	return e.msg
}
