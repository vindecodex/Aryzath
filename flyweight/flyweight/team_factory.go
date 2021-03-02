package flyweight

import "fmt"

const (
	TeamTypeA = "teamA"
	TeamTypeB = "teamB"
)

var (
	teamFactorySingleInstance = &TeamFactory{
		TeamMap: make(map[string]Team),
	}
)

type TeamFactory struct {
	TeamMap map[string]Team
}

func (t *TeamFactory) GetTeamType(teamType string) (Team, error) {
	if t.TeamMap[teamType] != nil {
		return t.TeamMap[teamType], nil
	}

	if teamType == TeamTypeA {
		t.TeamMap[teamType] = NewTeamA()
		return t.TeamMap[teamType], nil
	}

	if teamType == TeamTypeB {
		t.TeamMap[teamType] = NewTeamB()
		return t.TeamMap[teamType], nil
	}

	return nil, fmt.Errorf("Wrong team type passed")

}

func GetTeamFactorySingleInstance() *TeamFactory {
	return teamFactorySingleInstance
}
