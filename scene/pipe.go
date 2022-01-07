package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
	"math/rand"
)

var PipeImgUp *ebiten.Image
var PipeImgDown *ebiten.Image

func init() {
	pipeUpPng, _, err := image.Decode(bytes.NewReader(img.PipeUpPng))
	if err != nil {
		log.Fatal(err)
	}
	PipeImgUp = ebiten.NewImageFromImage(pipeUpPng)

	pipeDownPng, _, err := image.Decode(bytes.NewReader(img.PipeDownPng))
	if err != nil {
		log.Fatal(err)
	}
	PipeImgDown = ebiten.NewImageFromImage(pipeDownPng)
}

type Pipe struct {
	WidthUp  int
	HeightUp int

	WidthDown  int
	HeightDown int
}

func (b *Pipe) Init() {
	b.WidthUp, b.HeightUp = PipeImgUp.Size()
	b.WidthDown, b.HeightDown = PipeImgDown.Size()
}

func (b *Pipe) GapHigh() int {
	return rand.Intn(90)
}
