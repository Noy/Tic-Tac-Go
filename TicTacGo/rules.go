package TicTacGo

type Rules struct {}

func NewRules() *Rules { return new(Rules) }

func (r *Rules) MarkHasWinningSolution(board *GameBoard, mark string) bool {
	for _,set := range winningSets() {
		marks := marksForSet(board, set)
		if allMarksMatch(marks, mark) { return true }
	}
	return false
}

func marksForSet(board *GameBoard, set []int) []string {
	result := make([]string, len(set))
	for i,v := range set { result[i] = board.Spaces()[v] }
	return result
}

func allMarksMatch(marks []string, mark string) bool {
	result := true
	for _,v := range marks { result = result && v == mark }
	return result
}

func winningSets() [][]int {
  return [][]int{{0, 1, 2 }, {3, 4, 5 }, {6, 7, 8 },
		 {0, 3, 6 }, {1, 4, 7 }, {2, 5, 8 },
		 {0, 4, 8 }, {2, 4, 6 } }
}
