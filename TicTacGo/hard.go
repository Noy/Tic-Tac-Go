package TicTacGo

type Hard struct {
	mark string
	Minimax BoardScorer
}

type BoardScorer interface{
	ScoreAvailableMoves(*GameBoard, string) (map[int]int, bool)
	SetMinMaxMarks(string, string)
}

func LevelHard() *Hard {
	computer := new(Hard)
	computer.Minimax = NewTerm()
	return computer
}

func (c *Hard) Move(board GameBoard) int {
	moveScores,_ := c.Minimax.ScoreAvailableMoves(&board, c.mark)
	bestMove, bestScore := -1, -1
	for move,score := range moveScores {
		if score > bestScore { bestMove, bestScore = move, score }
	}
	return bestMove
}

func (c *Hard) SetMark(mark string) {
	c.mark = mark
	marks := map[string]string{ "X":"O", "O":"X" }
	c.Minimax.SetMinMaxMarks(marks[mark], mark)
}

func (c *Hard) GetMark() string { return c.mark}
