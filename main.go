package main

import (
	"image"
	"math"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Vec struct {
	X, Y float64
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {

	time := arange(0.0, 10.0, 0.1)
	angle := make([]float64, len(time))
	x := make([]float64, len(time))
	y := make([]float64, len(time))

	for i := 0; i < len(time); i++ {
		angle[i] = pendulum(time[i])
		x[i] = 0.5 * math.Sin(angle[i])
		y[i] = -0.5 * math.Cos(angle[i])
	}

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	pic, err := loadPicture("circle.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	win.Clear(colornames.Aliceblue)

	n := 0
	for !win.Closed() {
		if n > 99 {
			n = 0
		}
		var orig pixel.Vec
		orig.X = x[n] * 100
		orig.Y = y[n] * 100

		mat := pixel.IM
		mat = mat.ScaledXY(pixel.ZV, pixel.V(0.2, 0.2)).Moved(orig)
		sprite.Draw(win, mat)
		n++
		win.Update()
	}
}

func arange(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N)
	for x := range rnge {
		rnge[x] = start + step*float64(x)
	}
	return rnge
}

func pendulum(t float64) float64 {
	angleInit := math.Pi / 2
	natFrequency := math.Sqrt(9.8 / 0.5)
	angle := angleInit * math.Cos(natFrequency*t)
	return angle
}

func main() {
	pixelgl.Run(run)
}
