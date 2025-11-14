package domain

// GameStatus представляет состояние игры
type GameStatus struct {
	Finished          bool
	Win               bool
	Loss              bool
	CurrMistakes      int
	MaxMistakes       int
	State             string
	Word              string
	RemainingMistakes int
	Hint              string
	IsHintUsed        bool
}
