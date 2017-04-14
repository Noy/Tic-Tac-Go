package TicTacGo

type Game interface {
  Board() *GameBoard
  IsOver() bool
  IsValidMove(int) bool
  ApplyMove(int, string)
  Winner() (string, bool)
}

type game struct { board *GameBoard }

func NewGame() *game {
	g := new(game)
	g.board = NewBoard()
	return g
}

func (g *game) Board() *GameBoard { return g.board }

func (g *game) IsOver() bool { return winningFormatHappened(g.board) || full(g.board) }

func (g *game) Winner() (string, bool) {
	spaces := g.Board().Spaces()
	for _,set := range winningPositions() {
		if draw(g.Board(), set) {
			return spaces[ set[0] ], true
		}
	}
	return "", false
}

func (g *game) IsValidMove(space int) bool {
	board := g.board
	isInRange := space >= 0 && space < len(board.Spaces())
	return isInRange && board.Spaces()[ space ] == board.Blank()
}

func (g *game) ApplyMove(pos int, mark string) { if g.IsValidMove(pos) { g.board.Mark(pos, mark) } }

func (g *game) Reset() { g.board.Reset() }

func full(board *GameBoard) bool {
	for _,mark := range board.Spaces() {
		if mark == board.Blank() { return false }
	}
	return true
}

func winner(board *GameBoard) (string, bool) {
	for _,set := range winningPositions() {
		if draw(board, set) {
			return board.Spaces()[ set[0] ], true
		}
	}
	return "", false
}

func winningFormatHappened(board *GameBoard) (exists bool) {
	for _,set := range winningPositions() {
		exists = exists || draw(board, set)
	}
	return
}

func draw(board *GameBoard, pos []int) bool {
	spaces := board.Spaces()
	mark := spaces[ pos[ 0 ] ]
	result := mark != board.Blank()
	for _,i := range pos {
		result = result && spaces[ i ] == mark
	}
	return result
}

func winningPositions() [][]int {
	return [][]int{{0, 1, 2 }, {3, 4, 5 }, {6, 7, 8 },
		       {0, 3, 6 }, {1, 4, 7 }, {2, 5, 8 },
		       {0, 4, 8 }, {2, 4, 6 }}
}
