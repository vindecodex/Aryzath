package game

import (
	"Aryzath/flyweight/flyweight"
	"fmt"
)

type Game struct {
	Players []Player
}

func (g *Game) AddPlayer(team string) {
	player := newPlayer()
	teamFactoryInstance := flyweight.GetTeamFactorySingleInstance()
	playerTeam, _ := teamFactoryInstance.GetTeamType(team)
	player.setTeam(playerTeam)
	g.Players = append(g.Players, *player)
}

func (g *Game) ShowPlayers() {
	for _, player := range g.Players {
		fmt.Println(player.Team.GetColor())
	}
}
