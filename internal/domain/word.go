package domain

// Word представляет слово, используемое в игре "Виселица".
// Слово имеет значение, категорию и уровень сложности.
type Word struct {
	value      string
	category   string
	difficulty int
	hint       string
}

// NewWord возвращает новую структуру Word
func NewWord(value string, category string, difficulty int, hint string) Word {
	return Word{
		value:      value,
		category:   category,
		difficulty: difficulty,
		hint:       hint,
	}
}

// Value возвращает значение слова
func (w Word) Value() string {
	return w.value
}

// Category возвращает категорию слова
func (w Word) Category() string {
	return w.category
}

// Difficulty возвращает уровень сложности слова
func (w Word) Difficulty() int {
	return w.difficulty
}

func (w Word) Hint() string {
	return w.hint
}
