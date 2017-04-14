package TicTacGo

type Terms struct {
	DepthLimit int
	Rules *Rules
	currentDepth int
	minMark string
	maxMark string
}

func NewTerm() *Terms {
	m := new(Terms)
	m.Rules = NewRules()
	m.DepthLimit = 5
	return m
}

func (m *Terms) Score(board *GameBoard, currentMark string) (score int, final bool) {
	score, final = m.FinalScore(board)
	fin := false
	if !final && m.currentDepth < m.DepthLimit {
		m.currentDepth++
		nextMark := nextMark(currentMark)
		score,fin = m.bestOpposingScore(board, nextMark)
		final = final || fin || score == m.BestScoreForMark(nextMark)
		m.currentDepth--
	}
	return score, final
}

func (m *Terms) ScoreAvailableMoves(board *GameBoard, currentMark string) (map[int]int, bool) {
	availableSpaces := board.SpacesWithMark(board.Blank())
	result, final := make(map[int]int, len(availableSpaces)), true
	for _,space := range availableSpaces {
		board.Mark(space, currentMark)
		score,fin := m.Score(board, currentMark)
		result[space] = score
		final = final && fin
		board.Mark(space, board.Blank())
	}
	return result,final
}

func (m *Terms) FinalScore(board *GameBoard) (int, bool) {
	if m.Rules.MarkHasWinningSolution(board, m.maxMark) { return 1, true }
	if m.Rules.MarkHasWinningSolution(board, m.minMark) { return -1, true }
	if len(board.SpacesWithMark(board.Blank())) == 0 { return 0, true }
	return 0, false
}

func (m Terms) BestScoreForMark(mark string) int { if mark == m.maxMark { return 1 }; return -1 }

func (m *Terms) SetMinMaxMarks(min string, max string) {
	m.minMark = min
	m.maxMark = max
}

func (m *Terms) bestOpposingScore(board *GameBoard, nextMark string) (int,bool) {
	scores,fin := m.ScoreAvailableMoves(board, nextMark)
	return m.BestOfScores(scores, nextMark), fin
}

func nextMark(mark string) string { if mark == "X" { return "O" }; return "X" }

func (m *Terms) BestOfScores(scores map[int]int, mark string) int {
	target := m.BestScoreForMark(mark)
	best := target * -1
	for _,score := range scores {
		if (target == 1 && best < score) || (target == -1 && best > score) {
			best = score
		}
	}
	return best
}
