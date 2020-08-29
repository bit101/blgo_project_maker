package base

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

// Modes
const (
	ImageMode = iota
	GifMode
)

// Config contains the settings for creating images and animations.
type Config struct {
	Mode        int
	ImageWidth  float64
	ImageHeight float64
	GifWidth    float64
	GifHeight   float64
	Fps         float64
	Frames      int
	Render      func(*blgo.Surface, float64, float64, float64)
}

// Configure configures the drawing or animation.
func Configure(settings Config) {
	var surface *blgo.Surface
	var outName string

	switch settings.Mode {
	case ImageMode:
		outName = "out.png"
		surface = blgo.NewSurface(settings.ImageWidth, settings.ImageHeight)
		settings.Render(surface, settings.ImageWidth, settings.ImageHeight, 1.0)
		surface.WriteToPNG(outName)
		util.ViewImage(outName)

	case GifMode:
		outName = "out.gif"
		surface = blgo.NewSurface(settings.GifWidth, settings.GifHeight)
		anim := anim.NewAnimation(surface, settings.Frames, "frames")
		anim.Render(func(percent float64) {
			settings.Render(surface, settings.GifWidth, settings.GifHeight, percent)
		})
		util.ConvertToGIF("frames", outName, settings.Fps)
		util.ViewImage(outName)

	}

}
