package game

type Game struct {
	Version  string
	Title    string
	Theme    []string
	Platform string
}

func (g Game) Clone() *Game {
	return &Game{
		Version:  g.Version,
		Title:    g.Title,
		Theme:    g.Theme,
		Platform: g.Platform,
	}
}

func (g *Game) SetVersion(version string) {
	g.Version = version
}

func (g *Game) SetTitle(title string) {
	g.Title = title
}

func (g *Game) SetTheme(theme string) {
	g.Theme = append(g.Theme, theme)
}

func (g *Game) SetPlatform(platform string) {
	g.Platform = platform
}

func (g *Game) GetVersion() string {
	return g.Version
}

func (g *Game) GetTitle() string {
	return g.Title
}

func (g *Game) GetTheme() []string {
	return g.Theme
}

func (g *Game) GetPlatform() string {
	return g.Platform
}
