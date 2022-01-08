package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

var BirdImg [3]*ebiten.Image

func init() {
	birdPng0, _, err := image.Decode(bytes.NewReader(img.BirdPng0))
	if err != nil {
		log.Fatal(err)
	}
	birdPng1, _, err := image.Decode(bytes.NewReader(img.BirdPng1))
	if err != nil {
		log.Fatal(err)
	}
	birdPng2, _, err := image.Decode(bytes.NewReader(img.BirdPng2))
	if err != nil {
		log.Fatal(err)
	}
	BirdImg[0] = ebiten.NewImageFromImage(birdPng0)
	BirdImg[1] = ebiten.NewImageFromImage(birdPng1)
	BirdImg[2] = ebiten.NewImageFromImage(birdPng2)
}

const physicsCap = 10

type Birds struct {
	Width  int
	Height int

	WidthPhysics  int
	HeightPhysics int
}

func (b *Birds) Init() {
	b.Width, b.Height = BirdImg[0].Size()
	b.WidthPhysics, b.HeightPhysics = b.Width-physicsCap, b.Height-physicsCap
}
