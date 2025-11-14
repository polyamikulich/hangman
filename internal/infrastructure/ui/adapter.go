package ui

type UIAdapter struct{}

func NewUIAdapter() *UIAdapter {
	return &UIAdapter{}
}

func (uia *UIAdapter) DrawHangman(currMistakes, maxMistakes int) {
	DrawHangman(currMistakes, maxMistakes)
}

func (uia *UIAdapter) DrawHint(hint string) {
	DrawHint(hint)
}

func (uia *UIAdapter) DrawCurrentState(state string) {
	DrawCurrentState(state)
}

func (uia *UIAdapter) DrawAttempts(mistakes int) {
	DrawAttempts(mistakes)
}

func (uia *UIAdapter) ShowWin() {
	ShowWin()
}

func (uia *UIAdapter) ShowLoss(word string) {
	ShowLoss(word)
}
