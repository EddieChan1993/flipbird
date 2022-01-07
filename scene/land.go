package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

var LandImg *ebiten.Image

func init() {
	landPng, _, err := image.Decode(bytes.NewReader(img.LandPng))
	if err != nil {
		log.Fatal(err)
	}
	LandImg = ebiten.NewImageFromImage(landPng)
}

type Land struct {
	Width  int
	Height int
}

func (b *Land) Init() {
	b.Width, b.Height = LandImg.Size()
}
