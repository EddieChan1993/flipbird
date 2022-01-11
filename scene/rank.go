package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

var RankPanel *ebiten.Image

func init() {
	rankPanelPng, _, err := image.Decode(bytes.NewReader(img.RankPanelPng))
	if err != nil {
		log.Fatal(err)
	}
	RankPanel = ebiten.NewImageFromImage(rankPanelPng)
}

type Rank struct {
	RankPanelW int
	RankPanelH int
}

func (b *Rank) Init() {
	b.RankPanelW, b.RankPanelH = RankPanel.Size()
}
