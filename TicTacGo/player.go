package TicTacGo

type Player interface {
	Move(GameBoard) int
	SetMark( string )
	GetMark() string
}
