package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

var BgDayImg *ebiten.Image

func init() {
	bgDayPng, _, err := image.Decode(bytes.NewReader(img.BgDayPng))
	if err != nil {
		log.Fatal(err)
	}
	BgDayImg = ebiten.NewImageFromImage(bgDayPng)
}

type Bg struct {
	Width  int
	Height int
}

func (b *Bg) Init() {
	b.Width, b.Height = BgDayImg.Size()
}

