package src

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

//go:embed song.ogg
var songAudio []byte

type Game struct {
	Img *ebiten.Image
}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {

	if g.Img == nil {

		// Uncomment to play audio after Ebitengine initializes:
		// SetupGame()

		g.Img = ebiten.NewImage(32, 32)
		g.Img.Fill(color.RGBA{255, 0, 0, 255})
	}

	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(100, 100)
	screen.DrawImage(g.Img, opt)
}

func (g *Game) Layout(w, h int) (int, int) { return w, h }

func SetupGame() {

	fmt.Println("set window title")
	ebiten.SetWindowTitle("Android audio freeze")

	fmt.Println("create audio context")
	context := audio.NewContext(48000)

	fmt.Println("create bytes reader")
	byteReader := bytes.NewReader(songAudio)

	fmt.Println("decode ogg vorbis")
	decoded, err := vorbis.DecodeWithoutResampling(byteReader)

	if err != nil {
		panic(err)
	}

	fmt.Println("create player")
	player, err := context.NewPlayer(decoded)
	if err != nil {
		panic(err)
	}

	fmt.Println("player.Play()")
	player.Play()

	fmt.Println("game setup, audio playback done")

}
