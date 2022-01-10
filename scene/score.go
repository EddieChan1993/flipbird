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

const scorePath = "flappybird/fontscore"

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
}

type ScorePng struct {
	Width int
}

func (s *ScorePng) Init() {
	s.Width, _ = ScoreImg[0].Size()
}
func (s *ScorePng) ScoreDivide(score int) []int {
	if score <= 0 {
		return []int{0}
	}
	count := len(strconv.Itoa(score))
	base := int(math.Pow(10, float64(count-1)))
	res := make([]int, count)
	for i := 0; base != 0; i++ {
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
