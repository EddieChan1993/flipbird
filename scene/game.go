package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

var GameOverPng *ebiten.Image
var GameTitlePng *ebiten.Image
var GamePlayBtnPng *ebiten.Image
var GameRankBtnPng *ebiten.Image
var TutorialPng *ebiten.Image

func init() {
	gameOverPng, _, err := image.Decode(bytes.NewReader(img.GameOverPng))
	if err != nil {
		log.Fatal(err)
	}
	GameOverPng = ebiten.NewImageFromImage(gameOverPng)

	gameTitle, _, err := image.Decode(bytes.NewReader(img.TitlePng))
	if err != nil {
		log.Fatal(err)
	}
	GameTitlePng = ebiten.NewImageFromImage(gameTitle)

	gamePlayBtn, _, err := image.Decode(bytes.NewReader(img.PlayBtn))
	if err != nil {
		log.Fatal(err)
	}
	GamePlayBtnPng = ebiten.NewImageFromImage(gamePlayBtn)

	gameRankBtn, _, err := image.Decode(bytes.NewReader(img.RankBtn))
	if err != nil {
		log.Fatal(err)
	}
	GameRankBtnPng = ebiten.NewImageFromImage(gameRankBtn)

	tutorial, _, err := image.Decode(bytes.NewReader(img.Tutorial))
	if err != nil {
		log.Fatal(err)
	}
	TutorialPng = ebiten.NewImageFromImage(tutorial)
}

type GameExtra struct {
	GameOverTitleWidth  int
	GameStartTitleWidth int
	GamePlayBtnPngWidth int
	GameRankBtnPngWidth int
	TutorialWidth       int
	TutorialHeight      int
}

func (g *GameExtra) Init() {
	g.GameOverTitleWidth, _ = GameOverPng.Size()
	g.GameStartTitleWidth, _ = GameTitlePng.Size()
	g.GamePlayBtnPngWidth, _ = GamePlayBtnPng.Size()
	g.GameRankBtnPngWidth, _ = GameRankBtnPng.Size()
	g.TutorialWidth, g.TutorialHeight = TutorialPng.Size()
}
