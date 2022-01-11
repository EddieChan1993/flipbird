package radio

import (
	"embed"
	_ "embed"
)

//go:embed radio/fly.mp3
//go:embed radio/ding.mp3
//go:embed radio/hit.mp3
var gameRadio embed.FS

var FlyRadio []byte
var DingRadio []byte
var HitRadio []byte

func init() {
	var err error
	FlyRadio, err = gameRadio.ReadFile("radio/fly.mp3")
	if err != nil {
		panic(err)
	}
	DingRadio, err = gameRadio.ReadFile("radio/ding.mp3")
	if err != nil {
		panic(err)
	}
	HitRadio, err = gameRadio.ReadFile("radio/hit.mp3")
	if err != nil {
		panic(err)
	}
}
