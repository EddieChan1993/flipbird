package scene

import (
	"bytes"
	"flipbird/img"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io/fs"
	"log"
	"math"
	"strconv"
)

var ScoreImg [10]*ebiten.Image
var ScoreMiniImg [10]*ebiten.Image

const scorePath = "flappybird/fontscore"
const scoreMiniPath = "flappybird/numscore"

func init() {
	scorePngFiles, err := fs.ReadDir(img.ScorePngs, scorePath)
	if err != nil {
		log.Fatal(err)
	}
	//打印出文件名称
	for i, file := range scorePngFiles {
		scorePng, err := fs.ReadFile(img.ScorePngs, scorePath+"/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
		scoreP, _, err := image.Decode(bytes.NewReader(scorePng))
		if err != nil {
			log.Fatal(err)
		}
		ScoreImg[i] = ebiten.NewImageFromImage(scoreP)
	}

	scoreMiniPngFiles, err := fs.ReadDir(img.NumSmallPng, scoreMiniPath)
	if err != nil {
		log.Fatal(err)
	}
	//打印出文件名称
	for i, file := range scoreMiniPngFiles {
		scorePng, err := fs.ReadFile(img.NumSmallPng, scoreMiniPath+"/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
		scoreP, _, err := image.Decode(bytes.NewReader(scorePng))
		if err != nil {
			log.Fatal(err)
		}
		ScoreMiniImg[i] = ebiten.NewImageFromImage(scoreP)
	}
}

type ScorePng struct {
	ScoreWidth     int
	ScoreMiniWidth int
}

func (s *ScorePng) Init() {
	s.ScoreWidth, _ = ScoreImg[0].Size()
	s.ScoreMiniWidth, _ = ScoreMiniImg[0].Size()
}
func (s *ScorePng) ScoreDivide(score int) []int {
	if score <= 0 {
		return []int{0}
	}
	count := len(strconv.Itoa(score))
	base := int(math.Pow(10, float64(count-1)))
	res := make([]int, count)
	for i := count - 1; base != 0; i-- {
		b := score / base
		score = score - b*base
		if score <= 0 {
			score = 0
		}
		base /= 10
		res[i] = b
	}
	return res
}
