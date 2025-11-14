package input

type InputAdapter struct{}

func NewInputAdapter() *InputAdapter {
	return &InputAdapter{}
}

func (ia *InputAdapter) ChooseCategory() string {
	return ChooseCategory()
}

func (ia *InputAdapter) ChooseWordDifficulty() int {
	return ChooseWordDifficulty()
}

func (ia *InputAdapter) ChooseGameDifficulty() int {
	return ChooseGameDifficulty()
}

func (ia *InputAdapter) ChooseHint() int {
	return ChooseHint()
}

func (ia *InputAdapter) ReadNumber(maxnum int) (int, error) {
	return ReadNumber(maxnum)
}

func (ia *InputAdapter) ReadRune() rune {
	return ReadRune()
}
