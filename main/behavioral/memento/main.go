package main

func main() {
	game := &Game{hp: 100, mp: 100}

	memento := game.Save()
	game.Status()

	game.Play(-10, -20)
	game.Status()

	game.Load(&memento)
	game.Status()
}
