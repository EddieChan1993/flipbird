package img

import _ "embed"

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
