package main

import (
	"Aryzath/flyweight/flyweight"
	"Aryzath/flyweight/game"
)

func main() {
	game := &game.Game{}
	game.AddPlayer(flyweight.TeamTypeA)
	game.AddPlayer(flyweight.TeamTypeA)
	game.AddPlayer(flyweight.TeamTypeA)

	game.AddPlayer(flyweight.TeamTypeB)
	game.AddPlayer(flyweight.TeamTypeB)
	game.AddPlayer(flyweight.TeamTypeB)

	game.ShowPlayers()
}
