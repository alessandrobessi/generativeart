package main

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
	"image/color"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xFF, 0x75, 0xCC, 0xFF},
		{0x36, 0xC5, 0xFF, 0xFF},
		{0x77, 0xF7, 0x76, 0xFF},
		{0xFE, 0xFD, 0x32, 0xFF},
		{0xFE, 0x32, 0x72, 0xFF},
		{0xFF, 0xFF, 0xFF, 0xFF},
	}
	c := generativeart.NewCanva(2520, 3564)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewDotsWave(600))
	c.ToJPEG("no1.jpeg", 100)
}
