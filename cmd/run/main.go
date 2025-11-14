package main

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/data"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/input"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/infrastructure/ui"
)

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		inputReader := input.NewInputAdapter()
		wordRepo := data.NewWordRepositoryAdapter()
		uiInstance := ui.NewUIAdapter()

		if err := application.RunInteractiveMode(inputReader, wordRepo, uiInstance); err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}
	case 2:
		if err := application.RunTestMode(args[0], args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}
	default:
		application.ShowUsage()
		os.Exit(1)
	}
}
