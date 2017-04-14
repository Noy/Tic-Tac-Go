package TicTacGo

type UserPlayer struct {
  mark string
  console UI
}

func NewPlayer( console UI ) *UserPlayer {
	player := new(UserPlayer)
	player.console = console
	return player
}

func (c UserPlayer) Move( board GameBoard) int {
	available := board.SpacesWithMark(board.Blank())
	return c.console.PromptPlayerMove(available...)
}

func (c *UserPlayer) SetMark( mark string ) { c.mark = mark }

func (c UserPlayer) GetMark() string { return c.mark }
