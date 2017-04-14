package TicTacGo

type GameLogic struct {
	Game Game
	UI UI
	Players []Player
}

func (logic *GameLogic) Run() {
	selection := logic.UI.PromptMainMenu()
	if selection == EXIT_GAME { return }
	if selection == CONSOLE_FIRST {
		changePlayer(logic.Players)
	}
	runGameLoop(logic.Game, logic.Players, logic.UI)
}

func runGameLoop(game Game, players []Player, console UI) {
	current, next := players[0], players[1]
	for !game.IsOver() {
		console.DisplayAvailableSpaces(game.Board())
		move := current.Move(*game.Board())
		if game.IsValidMove(move) {
			game.ApplyMove(move, current.GetMark())
			current, next = next, current
		}
	}
	console.DisplayBoard(game.Board())
	console.DisplayGameOver(game)
}

func changePlayer(players []Player) {
	marks := []string{players[0].GetMark(), players[1].GetMark() }
	players[0], players[1] = players[1], players[0]
	players[0].SetMark(marks[0])
	players[1].SetMark(marks[1])
}
