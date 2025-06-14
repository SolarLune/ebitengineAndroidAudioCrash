package mobile

import (
	"ebitenAndroidAudioFreeze/src"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {

	src.SetupGame()

	fmt.Println("start game")

	mobile.SetGame(&src.Game{})
}

func Dummy() {}
