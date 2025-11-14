package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordRepositoryAdapter_Creation(t *testing.T) {
	repo := NewWordRepositoryAdapter()
	assert.NotNil(t, repo)
}

func TestWordRepositoryAdapter_GetRandomWord_Success(t *testing.T) {
	repo := NewWordRepositoryAdapter()

	// Пытаемся получить слово из существующей категории и сложности
	word, err := repo.GetRandomWord("животные", 1)

	assert.NoError(t, err)
	assert.NotEmpty(t, word.Value())
	assert.Equal(t, "животные", word.Category())
	assert.Equal(t, 1, word.Difficulty())
}

func TestGetRandomWord_WithHint(t *testing.T) {
	// Пытаемся получить слово из существующей категории и сложности
	word, err := GetRandomWord("страны", 2)

	assert.NoError(t, err)
	assert.NotEmpty(t, word.Value())
	assert.Equal(t, "страны", word.Category())
	assert.Equal(t, 2, word.Difficulty())
	assert.NotEmpty(t, word.Hint(), "У полученного слова должна быть подсказка")
}

func TestWordRepositoryAdapter_GetRandomWord_CategoryNotFound(t *testing.T) {
	repo := NewWordRepositoryAdapter()

	// Пытаемся получить слово из несуществующей категории
	_, err := repo.GetRandomWord("несуществующая_категория", 1)

	assert.Error(t, err)
	assert.Equal(t, ErrNoWordsFound, err)
}

func TestWordRepositoryAdapter_GetRandomWord_DifficultyNotFound(t *testing.T) {
	repo := NewWordRepositoryAdapter()

	// Пытаемся получить слово с недопустимой сложностью в существующей категории
	_, err := repo.GetRandomWord("животные", 5)

	assert.Error(t, err)
	assert.Equal(t, ErrNoWordsFound, err)
}

func TestWordRepositoryAdapter_GetRandomWord_InvalidCategoryAndDifficulty(t *testing.T) {
	repo := NewWordRepositoryAdapter()

	// Пытаемся получить слово с недопустимыми параметрами
	_, err := repo.GetRandomWord("несуществующая_категория", 10)

	assert.Error(t, err)
	assert.Equal(t, ErrNoWordsFound, err)
}

func TestDictionaryStructure(t *testing.T) {
	// Проверяем, что словарь не пустой
	assert.NotEmpty(t, Dictionary)

	// Проверяем, что все слова имеют допустимую сложность (1, 2 или 3)
	validDifficulties := []int{1, 2, 3}

	for _, word := range Dictionary {
		assert.Contains(t, validDifficulties, word.Difficulty())
		assert.NotEmpty(t, word.Value())
		assert.NotEmpty(t, word.Category())
	}
}

func TestDictionary_CategoriesExist(t *testing.T) {
	// Проверяем, что в словаре есть слова из всех ожидаемых категорий
	expectedCategories := []string{"животные", "фрукты", "страны", "профессии", "музыкальные инструменты"}

	foundCategories := make(map[string]bool)
	for _, category := range expectedCategories {
		foundCategories[category] = false
	}

	for _, word := range Dictionary {
		if _, exists := foundCategories[word.Category()]; exists {
			foundCategories[word.Category()] = true
		}
	}

	for category, found := range foundCategories {
		assert.True(t, found, "Категория %s должна присутствовать в словаре", category)
	}
}

func TestDictionary_DifficultyDistribution(t *testing.T) {
	// Проверяем, что в каждой категории есть слова всех уровней сложности
	categories := make(map[string]map[int]int)

	for _, word := range Dictionary {
		if _, exists := categories[word.Category()]; !exists {
			categories[word.Category()] = make(map[int]int)
		}
		categories[word.Category()][word.Difficulty()]++
	}

	expectedDifficulties := []int{1, 2, 3}

	for category, difficulties := range categories {
		for _, difficulty := range expectedDifficulties {
			count, exists := difficulties[difficulty]
			assert.True(t, exists, "В категории %s должны быть слова сложности %d", category, difficulty)
			assert.Greater(t, count, 0, "В категории %s должны быть слова сложности %d", category, difficulty)
		}
	}
}

func TestDictionary_Hints(t *testing.T) {
	// Проверяем, что у всех слов в словаре есть  непустые подсказки
	for _, word := range Dictionary {
		assert.NotEmpty(t, word.Hint(), "У слова '%s' из категории '%s' должна быть подсказка", word.Value(), word.Category())
	}
}

func TestErrNoWordsFound_Error(t *testing.T) {
	// Проверяем, что ошибка ErrNoWordsFound определена правильно
	assert.EqualError(t, ErrNoWordsFound, "No words found for the given category and difficulty")
}
