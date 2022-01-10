package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestAa(t *testing.T) {
	score := int(0)
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
	fmt.Println(res)
}
