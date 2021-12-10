package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"
	"os"
)

const ImgWidth = 600
const ImgHeight = 600
const CandleWidth = 100

func CreateGif(out io.Writer) {

	imgs, lastRect := displayCandle(100, 100, 200, []image.Rectangle{})
	imgs2, _ := displayCandle(250, 150, 250, []image.Rectangle{lastRect})
	imgs = append(imgs, imgs2...)
	delays := make([]int, len(imgs))
	for i := 0; i < len(imgs); i++ {
		delays[i] = 100 * 4 / len(imgs)
	}

	anim := gif.GIF{Image: imgs, Delay: delays}

	gif.EncodeAll(out, &anim)
}

func displayCandle(x, y, height int, keepRectangles []image.Rectangle) ([]*image.Paletted, image.Rectangle) {

	var images []*image.Paletted
	var redRect image.Rectangle
	step := height / 60

	for i := 0; i < 60; i++ {
		myred := color.RGBA{200, 0, 0, 255}
		palette := []color.Color{color.White, color.Black, myred}

		rect := image.Rect(0, 0, ImgWidth, ImgHeight)
		img := image.NewPaletted(rect, palette)

		redRect = image.Rect(x, y, x+CandleWidth, y+i*step) //  geometry of 2nd rectangle which we draw atop above rectangle

		draw.Draw(img, redRect, &image.Uniform{myred}, image.ZP, draw.Src)
		for _, keepRect := range keepRectangles {
			draw.Draw(img, keepRect, &image.Uniform{myred}, image.ZP, draw.Src)
		}

		images = append(images, img)
	}

	return images, redRect
}

func main() {
	f, err := os.Create("my-image.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	CreateGif(f)
}
