package data

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw1-hangman/internal/domain"
)

type WordRepositoryAdapter struct{}

func NewWordRepositoryAdapter() *WordRepositoryAdapter {
	return &WordRepositoryAdapter{}
}

func (wra *WordRepositoryAdapter) GetRandomWord(category string, difficulty int) (domain.Word, error) {
	return GetRandomWord(category, difficulty)
}
