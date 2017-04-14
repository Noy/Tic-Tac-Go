package main

import (
        "os"
         tictacgo "./TicTacGo"
)


func main() { start() }

func start() {
	console := tictacgo.NewConsole(os.Stdin, os.Stdout)
	game := tictacgo.NewGame()

	player1 := tictacgo.NewPlayer(console)
	// You can actually make them play each other, it's funny to watch haha
	// consolePlayer2 := tictacgo.LevelLEasy()
	consolePlayer := tictacgo.LevelHard()

	players := []tictacgo.Player{player1, consolePlayer}

	player1.SetMark("x")
	consolePlayer.SetMark("o") // You can actually use any symbol you want

	logic := tictacgo.GameLogic{Game: game, UI: console, Players: players }
	logic.Run()
}
