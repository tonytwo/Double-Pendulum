package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	var pendulums []sPendulum
	pendulums = append(pendulums, sPendulum{
		grav:      1.5,
		radius1:   400.0,
		radius2:   400.0,
		mass1:     64.0,
		mass2:     64.0,
		thetaVel1: 0.0,
		thetaVel2: 0.0,
		theta1:    3.1415 / 1.2,
		theta2:    3.1415 / 2.0,
		color:     color.NRGBA{255, 0, 0, 255},
	},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1414 / 1.2,
			theta2:    3.1414 / 2.0,
			color:     color.NRGBA{255, 0, 50, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1413 / 1.2,
			theta2:    3.1413 / 2.0,
			color:     color.NRGBA{255, 0, 100, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1412 / 1.2,
			theta2:    3.1412 / 2.0,
			color:     color.NRGBA{255, 0, 150, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1411 / 1.2,
			theta2:    3.1411 / 2.0,
			color:     color.NRGBA{255, 0, 200, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1410 / 1.2,
			theta2:    3.1410 / 2.0,
			color:     color.NRGBA{255, 0, 250, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1409 / 1.2,
			theta2:    3.1409 / 2.0,
			color:     color.NRGBA{200, 0, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1408 / 1.2,
			theta2:    3.1408 / 2.0,
			color:     color.NRGBA{150, 0, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1407 / 1.2,
			theta2:    3.1407 / 2.0,
			color:     color.NRGBA{100, 0, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1406 / 1.2,
			theta2:    3.1406 / 2.0,
			color:     color.NRGBA{50, 0, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1405 / 1.2,
			theta2:    3.1405 / 2.0,
			color:     color.NRGBA{0, 0, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1404 / 1.2,
			theta2:    3.1404 / 2.0,
			color:     color.NRGBA{0, 50, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1403 / 1.2,
			theta2:    3.1403 / 2.0,
			color:     color.NRGBA{0, 100, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1402 / 1.2,
			theta2:    3.1402 / 2.0,
			color:     color.NRGBA{0, 150, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1401 / 1.2,
			theta2:    3.1401 / 2.0,
			color:     color.NRGBA{0, 200, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1400 / 1.2,
			theta2:    3.1400 / 2.0,
			color:     color.NRGBA{0, 255, 255, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1399 / 1.2,
			theta2:    3.1399 / 2.0,
			color:     color.NRGBA{0, 255, 200, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1398 / 1.2,
			theta2:    3.1398 / 2.0,
			color:     color.NRGBA{0, 255, 150, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1397 / 1.2,
			theta2:    3.1397 / 2.0,
			color:     color.NRGBA{0, 255, 100, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1396 / 1.2,
			theta2:    3.1396 / 2.0,
			color:     color.NRGBA{0, 255, 50, 255},
		},
		sPendulum{
			grav:      1.5,
			radius1:   400.0,
			radius2:   400.0,
			mass1:     64.0,
			mass2:     64.0,
			thetaVel1: 0.0,
			thetaVel2: 0.0,
			theta1:    3.1395 / 1.2,
			theta2:    3.1395 / 2.0,
			color:     color.NRGBA{0, 255, 0, 255},
		})

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

		win.Clear(color.NRGBA{44, 44, 84, 255})
		for i := 0; i < len(pendulums); i++ {
			pendulums[i].draw(win)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
