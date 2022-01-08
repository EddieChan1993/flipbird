package main

import (
	"flipbird/scene"
	_ "flipbird/scene"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	_ "image/png"
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

	pipeStartOffsetX = 1
	pipeGapY         = 90
	pipeGapX         = 150
)

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

var ()

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
	bg    scene.Bg
	land  scene.Land
	pipe  scene.Pipe
	birds scene.Birds
	mode  Mode

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
	g.birds.Init()
	g.x16 = 0
	g.y16 = 4000
	g.cameraX = -150
	g.cameraY = 0
	g.count = 0
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
		if g.hit() {
			fmt.Println("碰到了")
			//g.hitPlayer.Rewind()
			//g.hitPlayer.Play()
			g.mode = ModeGameOver
			g.gameoverCount = 30
		}
	case ModeGameOver:
		if g.gameoverCount > 0 {
			g.gameoverCount--
		}
		if g.gameoverCount == 0 && g.isKeyJustPressed() {
			fmt.Println("游戏结束")
			g.init()
			g.mode = ModeTitle
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawScene(screen)
	g.drawBird(screen)
	g.drawPipe(screen)
}

func (g *Game) drawBird(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	index := (g.count / 5) % 3
	w, h := g.birds.Width, g.birds.Height
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Rotate(float64(g.vy16) / 96.0 * math.Pi / 6)
	op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)
	op.GeoM.Translate(float64(g.x16/16.0)-float64(g.cameraX), float64(g.y16/16.0)-float64(g.cameraY))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(scene.BirdImg[index], op)
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
	pipeIndex := 1
	birthW, birthH := g.birds.WidthPhysics, g.birds.HeightPhysics
	x0 := float64(g.x16/16.0) - float64(g.cameraX)
	y0 := float64(g.y16/16.0) - float64(g.cameraY)
	y1 := y0 + float64(birthH)
	//上限
	if y0 < 0 {
		return true
	}
	//下限
	if y1 >= float64(screenHeight-g.land.Height) {
		return true
	}
	xMin := floorDiv(g.count, g.pipe.WidthDown+pipeGapX)
	pipeGapHigh, ok := g.pipeAt(xMin + pipeIndex)
	if ok {
		//最近的管子起始坐标
		pipeX := float64(pipeIndex*(g.pipe.WidthDown+pipeGapX) - floorMod(g.count, g.pipe.WidthDown+pipeGapX))
		pipeY := -70 - pipeGapY + float64(pipeGapHigh) + float64(g.pipe.HeightUp)
		if pipeX <= x0+float64(birthW) && pipeX >= x0-float64(g.pipe.WidthUp) {
			if y0 <= pipeY || y1 >= pipeY+float64(pipeGapY+70) {
				return true
			}
		}
	}
	return false
}

func (g *Game) drawPipe(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < 3; i++ {
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
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flappy Gopher (Ebiten Demo)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
