package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var (
	grav           = 1.0
	radius1        = 300.0
	radius2        = 400.0
	mass1          = 64.0
	mass2          = 64.0
	thetaVel1      = 0.0
	thetaVel2      = 0.0
	theta1, theta2 = a1a2DefaultValues()
)

func run() {
	winWidth := 1080 * 2
	winHeight := 786 * 2
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, float64(winWidth), float64(winHeight)),
		VSync:       true,
		Undecorated: true,
	})
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	win.SetMatrix(pixel.IM.ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(pixel.V(float64(winWidth)/2, float64(winHeight)*3/4)))

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		if win.JustPressed(pixelgl.KeySpace) {
			theta1, theta2 = a1a2DefaultValues()
		}

		win.Clear(color.NRGBA{44, 44, 84, 255})

		a, b := update()

		imd := imdraw.New(nil)

		imd.Color = color.NRGBA{64, 64, 122, 255}
		imd.Push(pixel.ZV, a, b)
		imd.Line(3)

		imd.Color = color.NRGBA{51, 217, 178, 255}
		imd.Push(a)
		imd.Circle(mass1/2, 0)

		imd.Color = color.NRGBA{52, 172, 224, 255}
		imd.Push(b)
		imd.Circle(mass2/2, 0)

		imd.Draw(win)
		win.Update()
	}
}

func update() (pixel.Vec, pixel.Vec) {
	a1a := a1aCalculation()
	a2a := a2aCalculation()

	thetaVel1 += a1a
	thetaVel2 += a2a

	theta1 += thetaVel1
	theta2 += thetaVel2

	thetaVel1 *= 0.9996
	thetaVel2 *= 0.9996

	a := pixel.V(radius1*math.Sin(theta1), radius1*math.Cos(theta1))
	b := pixel.V(a.X+radius2*math.Sin(theta2), a.Y+radius2*math.Cos(theta2))

	return a, b
}

func main() {
	pixelgl.Run(run)
}

func a1a2DefaultValues() (float64, float64) {
	return math.Pi / 2, math.Pi / 2
}

func a1aCalculation() float64 {
	num1 := -grav * (2*mass1 + mass2) * math.Sin(theta1)
	num2 := -mass2 * grav * math.Sin(theta1-2*theta2)
	num3 := -2 * math.Sin(theta1-theta2) * mass2
	num4 := thetaVel2*thetaVel2*radius2 + thetaVel1*thetaVel1*radius1*math.Cos(theta1-theta2)
	den := radius1 * (2*mass1 + mass2 - mass2*math.Cos(2*theta1-2*theta2))

	return (num1 + num2 + num3*num4) / den
}

func a2aCalculation() float64 {
	num1 := 2 * math.Sin(theta1-theta2)
	num2 := (thetaVel1 * thetaVel1 * radius1 * (mass1 + mass2))
	num3 := grav * (mass1 + mass2) * math.Cos(theta1)
	num4 := thetaVel2 * thetaVel2 * radius2 * mass2 * math.Cos(theta1-theta2)
	den := radius2 * (2*mass1 + mass2 - mass2*math.Cos(2*theta1-2*theta2))

	return (num1 * (num2 + num3 + num4)) / den
}
