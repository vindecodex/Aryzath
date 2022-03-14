package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/prototype/game"
)

func main() {
	// Version 1 of our game only has a theme of Action and war
	theme := []string{"Action", "War"}
	game_v1 := game.Game{
		Version:  "v1.0.0",
		Title:    "Game of Thrones",
		Theme:    theme,
		Platform: "PC",
	}

	// Version 2 of our game added an Adventure theme
	game_v2 := game_v1.Clone()
	game_v2.SetVersion("v2.0.0")
	game_v2.SetTheme("Adventure")

	fmt.Println("Game V1: ", game_v1)  // version 1 of our game
	fmt.Println("Game V2: ", *game_v2) // version 2 of our game

	// We can create another version using game_v2 or game_v1

	game_v3 := game_v2.Clone()
	game_v4 := game_v2.Clone()

	// No changes made
	fmt.Println("Game V3: ", game_v3)
	fmt.Println("Game v4: ", game_v4)
}
