package avatars

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type GeneratorRules struct {
}

type ImageGenerator struct {
}

func (ig *ImageGenerator) NewGenerationConfig() GenerationConfig {
	return GenerationConfig{}
}

// Describes expected generated image result
type GenerationConfig struct {
	Width           int // Image width
	Height          int // Image height
	BackgroundColor color.Color
	PixelColor      color.Color
	Pixels          [][]int
	PixelSize       uint16
	PixelShape      int
	ImageShape      int
}

func (ig *ImageGenerator) CreateCanvas(conf *GenerationConfig) *image.RGBA {

	canvas := image.NewRGBA(image.Rect(0, 0, conf.Width, conf.Height))
	draw.Draw(
		canvas,
		canvas.Bounds(),
		image.NewUniform(conf.BackgroundColor),
		image.Point{},
		draw.Src,
	)

	return canvas
}

func (ig *ImageGenerator) GrawImage(conf *GenerationConfig, canvas *image.RGBA) *image.RGBA {
	draw.Draw(
		canvas,
		image.Rect(0, 0, 50, 50),
		image.NewUniform(conf.PixelColor),
		image.Point{25, 25},
		draw.Over,
	)
	return canvas
}

func (ig *ImageGenerator) ExportPNG(result *image.RGBA) []byte {
	buffer := bytes.NewBuffer(nil)
	png.Encode(buffer, result)
	return buffer.Bytes()
}

func (ig *ImageGenerator) CreateImage() {

	conf := GenerationConfig{
		Width:           100,
		Height:          100,
		BackgroundColor: image.NewUniform(color.White),
		PixelColor:      image.NewUniform(color.Black),
	}

	canvas := ig.CreateCanvas(&conf)
	ig.GrawImage(&conf, canvas)

	os.WriteFile("out.png", ig.ExportPNG(canvas), os.ModePerm)
}
