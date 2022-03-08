package game

import "github.com/vindecodex/Aryzath/flyweight/flyweight"

type Player struct {
	Team flyweight.Team
}

func (p *Player) getTeamColor() string {
	return p.Team.GetColor()
}

func (p *Player) setTeam(team flyweight.Team) {
	p.Team = team
}

func newPlayer() *Player {
	return &Player{}
}
