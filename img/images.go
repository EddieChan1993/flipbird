package img

import (
	"embed"
	_ "embed"
)

//go:embed flappybird/bird0_0.png
var BirdPng0 []byte

//go:embed flappybird/bird0_1.png
var BirdPng1 []byte

//go:embed flappybird/bird0_2.png
var BirdPng2 []byte

//go:embed flappybird/land.png
var LandPng []byte

//go:embed flappybird/bg_day.png
var BgDayPng []byte

//go:embed flappybird/pipe_up.png
var PipeUpPng []byte

//go:embed flappybird/pipe_down.png
var PipeDownPng []byte

//go:embed flappybird/fontscore/*
var ScorePngs embed.FS

//go:embed flappybird/text_game_over.png
var GameOverPng []byte

//go:embed flappybird/title.png
var TitlePng []byte

//go:embed flappybird/button_play.png
var PlayBtn []byte

//go:embed flappybird/button_score.png
var RankBtn []byte

//go:embed flappybird/tutorial.png
var Tutorial []byte

//go:embed flappybird/score_panel.png
var RankPanelPng []byte

//go:embed flappybird/numscore/*
var NumSmallPng embed.FS
