package TicTacGo

type GameBoard struct { spaces []string }

func (b GameBoard) Blank() string { return " " }

func (b *GameBoard) Spaces() []string {
	dup := make([]string, len(b.spaces))
	copy(dup, b.spaces)
	return dup
}

func (b *GameBoard) Mark(pos int, mark string) { if pos >= 0 && pos < len(b.spaces) { b.spaces[ pos ] = mark } }

func (b GameBoard) SpacesWithMark(mark string) []int {
	count, result := 0, make([]int, len(b.Spaces()))
	for i, v := range b.Spaces() {
		if v == mark {
			result[count] = i
			count++
		}
	}
	return result[:count]
}

func (b *GameBoard) Reset() { setBoard(b) }

func NewBoard() *GameBoard {
	b := new(GameBoard)
	setBoard(b)
	return b
}

func setBoard(b *GameBoard) {
	b.spaces = make([]string, 9)
	for i := range b.spaces { b.spaces[ i ] = " " }
}
