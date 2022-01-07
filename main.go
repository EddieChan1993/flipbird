package main

import (
	"bytes"
	"flipbird/img"
	"flipbird/scene"
	_ "flipbird/scene"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 400
	screenHeight = 600
	tileSize     = 32

	pipeStartOffsetX = 1
	pipeGapY         = 90
	pipeGapX         = 150
	pipeWidth        = 52
)

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

var (
	bird [3]*ebiten.Image
)

type Mode int

func floorDiv(x, y int) int {
	d := x / y
	if d*y == x || x >= 0 {
		return d
	}
	return d - 1
}

func floorMod(x, y int) int {
	return x - floorDiv(x, y)*y
}

func (g *Game) pipeAt(tileX int) (tileY int, ok bool) {
	if (tileX - pipeStartOffsetX) <= 0 {
		return 0, false
	}
	//if floorMod(tileX-pipeStartOffsetX, pipeGapX) != 0 {
	//	return 0, false
	//}
	//idx := floorDiv(tileX-pipeStartOffsetX, pipeGapX)
	return g.pipeTileYs[tileX%len(g.pipeTileYs)], true
}

type Game struct {
	bg   scene.Bg
	land scene.Land
	pipe scene.Pipe
	mode Mode

	count int

	// The bird's position
	x16  int
	y16  int
	vy16 int

	// Camera
	cameraX int
	cameraY int

	// Pipes
	pipeTileYs    []int
	gameoverCount int
}

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
	bird[0] = ebiten.NewImageFromImage(birdPng0)
	bird[1] = ebiten.NewImageFromImage(birdPng1)
	bird[2] = ebiten.NewImageFromImage(birdPng2)
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.bg.Init()
	g.land.Init()
	g.pipe.Init()
	g.x16 = 0
	g.y16 = 100 * 16
	g.cameraX = -150
	g.cameraY = 0
	g.pipeTileYs = make([]int, 256)
	for i := range g.pipeTileYs {
		g.pipeTileYs[i] = rand.Intn(pipeGapY)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.isKeyJustPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		g.count++
		g.x16 += 32
		g.cameraX += 2
		if g.isKeyJustPressed() {
			g.vy16 = -96
			//g.jumpPlayer.Rewind()
			//g.jumpPlayer.Play()
		}
		g.y16 += g.vy16

		// Gravity
		g.vy16 += 4
		if g.vy16 > 96 {
			g.vy16 = 96
		}
		//if g.hit() {
		//	fmt.Println("碰到了")
		//	//g.hitPlayer.Rewind()
		//	//g.hitPlayer.Play()
		//	g.mode = ModeGameOver
		//	g.gameoverCount = 30
		//}
	case ModeGameOver:
		if g.gameoverCount > 0 {
			g.gameoverCount--
		}
		if g.gameoverCount == 0 && g.isKeyJustPressed() {
			g.init()
			g.mode = ModeTitle
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//panic("implement me")
	g.drawScene(screen)
	g.drawBird(screen)
	//w, h := bird[0].Size()
	var (
		gopherWidth  = 48
		gopherHeight = 48
	)

	x0 := float64(g.x16/16.0) - float64(g.cameraX)
	y0 := float64(g.y16/16.0) - float64(g.cameraY)
	//fmt.Println(x0, y0)
	//x1 := x0 + float64(gopherWidth)
	//y1 := y0 + float64(gopherHeight)
	ebitenutil.DrawRect(screen, float64(x0), float64(y0), float64(gopherWidth), float64(gopherHeight), color.RGBA{255, 0, 0, 255})

	xMin := floorDiv(g.count, g.pipe.WidthDown+pipeGapX)
	pipeGapHigh, ok := g.pipeAt(xMin)
	if ok {
		fmt.Println("----------------ok----", pipeGapHigh)
		y := -70 - pipeGapY + float64(pipeGapHigh)
		ebitenutil.DrawRect(screen, float64(0*(g.pipe.WidthDown+pipeGapX)-floorMod(g.count, g.pipe.WidthDown+pipeGapX)), float64(y), float64(pipeWidth), float64(gopherHeight), color.RGBA{255, 0, 0, 255})
	}
}

func (g *Game) drawBird(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	index := (g.count / 5) % 3
	w, h := bird[index].Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Rotate(float64(g.vy16) / 96.0 * math.Pi / 6)
	op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)
	op.GeoM.Translate(float64(g.x16/16.0)-float64(g.cameraX), float64(g.y16/16.0)-float64(g.cameraY))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(bird[index], op)
}

func (g *Game) isKeyJustPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}

func (g *Game) hit() bool {
	if g.mode != ModeGame {
		return false
	}
	const (
		gopherWidth  = 20
		gopherHeight = 40
	)
	//w, h := bird[0].Size()
	x0 := int(float64(g.x16/16.0) - float64(g.cameraX))
	y0 := int(float64(g.y16/16.0) - float64(g.cameraY))
	xMin := floorDiv(x0-pipeWidth, tileSize)
	fmt.Println("xMin-----", xMin, "g.cameraX-----", g.cameraX, "x-----", x0, "y-----", y0)
	//x1 := x0 + gopherWidth
	y1 := y0 + gopherHeight
	//上限
	if y0 < -tileSize*4 {
		return true
	}
	//下限
	if y1 >= screenHeight-112 {
		return true
	}
	return false
}

func (g *Game) drawScene(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < 3; i++ {
		//背景
		op.GeoM.Reset()
		op.GeoM.Translate(float64(i*g.bg.Width-floorMod(g.count, g.bg.Width)), -24)
		screen.DrawImage(scene.BgDayImg, op)
		//地面
		op.GeoM.Reset()
		op.GeoM.Translate(float64(i*g.land.Width-floorMod(g.count, g.land.Width)), float64(screenHeight-g.land.Height))
		screen.DrawImage(scene.LandImg, op)
		//障碍物
		if pipeGapHigh, ok := g.pipeAt(floorDiv(g.count, g.pipe.WidthDown+pipeGapX) + i); ok {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(i*(g.pipe.WidthDown+pipeGapX)-floorMod(g.count, g.pipe.WidthDown+pipeGapX)), -70-pipeGapY+float64(pipeGapHigh))
			screen.DrawImage(scene.PipeImgDown, op)

			op.GeoM.Reset()
			op.GeoM.Translate(float64(i*(g.pipe.WidthDown+pipeGapX)-floorMod(g.count, g.pipe.WidthDown+pipeGapX)), screenHeight-370+pipeGapY+float64(pipeGapHigh))
			screen.DrawImage(scene.PipeImgUp.SubImage(image.Rect(0, 0, g.pipe.WidthDown, 320-112-40-pipeGapHigh)).(*ebiten.Image), op)
		}
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flappy Gopher (Ebiten Demo)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
