package main

import (
	"ebitenAndroidAudioFreeze/src"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// Call SetupGame() before Ebitengine's loop starts
	src.SetupGame()

	fmt.Println("start game")

	if err := ebiten.RunGame(&src.Game{}); err != nil {
		panic(err)
	}

}
