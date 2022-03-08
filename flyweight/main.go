package main

import (
	"github.com/vindecodex/Aryzath/flyweight/flyweight"
	"github.com/vindecodex/Aryzath/flyweight/game"
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
