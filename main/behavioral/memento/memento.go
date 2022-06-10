package main

import "fmt"

type Memento interface {
}

type GameMemento struct {
	hp, mp int
}

type Game struct {
	hp, mp int
}

func (game *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", game.hp, game.mp)
}

func (game *Game) Play(mpDelta, hpDelta int) {
	game.mp += mpDelta
	game.hp += hpDelta
}

func (game *Game) Save() Memento {
	return &GameMemento{
		hp: game.hp,
		mp: game.mp,
	}
}

func (game *Game) Load(memento *Memento) {
	gameGameMemento := (*memento).(*GameMemento)
	game.hp = (*gameGameMemento).hp
	game.mp = (*gameGameMemento).mp
}
