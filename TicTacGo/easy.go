package TicTacGo

import "math/rand"

type Easy struct { mark string }

func LevelEasy() *Easy { return new(Easy) }

func (e *Easy) Move( b GameBoard) (move int) {
	for move = rand.Intn(9); b.Spaces()[move] != b.Blank(); {
		move = rand.Intn(9)
	}
	return
}

func (e *Easy) SetMark( mark string ) { e.mark = mark }

func (e *Easy) GetMark() string { return e.mark }
