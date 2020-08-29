package main

import (
	"sketch/base"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
)

var settings = base.Config{
	Mode:        base.ImageMode,
	ImageWidth:  1920,
	ImageHeight: 1080,
	GifWidth:    400,
	GifHeight:   400,
	Fps:         30,
	Frames:      30 * 2,
	Render:      Render,
}

func main() {
	base.Configure(settings)
}

// Render renders a single frame.
func Render(surface *blgo.Surface, width, height, percent float64) {
	surface.ClearRGB(1, 1, 1)

	surface.StrokeCircle(width/2, height/2, width/2*blmath.LerpSin(percent, 1, 0))

}
