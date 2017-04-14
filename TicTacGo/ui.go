package TicTacGo

const (
	EXIT_GAME = iota
	PLAYER_FIRST
	CONSOLE_FIRST
)

type UI interface {
	PromptMainMenu() int
	DisplayBoard(*GameBoard)
	PromptPlayerMove(...int) int
	DisplayAvailableSpaces(*GameBoard)
	DisplayGameOver(Game)
}

type Reader interface { Read(b []byte) (n int, err error) }

type Writer interface { WriteString(s string) (ret int, err error) }
