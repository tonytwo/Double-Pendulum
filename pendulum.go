package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type sPendulum struct {
	grav      float64
	radius1   float64
	radius2   float64
	mass1     float64
	mass2     float64
	thetaVel1 float64
	thetaVel2 float64
	theta1    float64
	theta2    float64
	pendPos1  pixel.Vec
	pendPos2  pixel.Vec
	color     color.NRGBA
}

func (pendulum *sPendulum) update() {
	pendulum.a1aCalculation()
	pendulum.a2aCalculation()

	pendulum.theta1 += pendulum.thetaVel1
	pendulum.theta2 += pendulum.thetaVel2

	pendulum.thetaVel1 *= 1
	pendulum.thetaVel2 *= 1

	pendulum.pendPos1 = pixel.V(pendulum.radius1*math.Sin(pendulum.theta1), pendulum.radius1*math.Cos(pendulum.theta1))
	pendulum.pendPos2 = pixel.V(pendulum.pendPos1.X+pendulum.radius2*math.Sin(pendulum.theta2), pendulum.pendPos1.Y+pendulum.radius2*math.Cos(pendulum.theta2))
}

func (pendulum *sPendulum) a1aCalculation() {
	num1 := -pendulum.grav * (2*pendulum.mass1 + pendulum.mass2) * math.Sin(pendulum.theta1)
	num2 := -pendulum.mass2 * pendulum.grav * math.Sin(pendulum.theta1-2*pendulum.theta2)
	num3 := -2 * math.Sin(pendulum.theta1-pendulum.theta2) * pendulum.mass2
	num4 := pendulum.thetaVel2*pendulum.thetaVel2*pendulum.radius2 + pendulum.thetaVel1*pendulum.thetaVel1*pendulum.radius1*math.Cos(pendulum.theta1-pendulum.theta2)
	den := pendulum.radius1 * (2*pendulum.mass1 + pendulum.mass2 - pendulum.mass2*math.Cos(2*pendulum.theta1-2*pendulum.theta2))

	pendulum.thetaVel1 += (num1 + num2 + num3*num4) / den
}

func (pendulum *sPendulum) draw(win pixel.Target) {
	pendulum.update()

	imd := imdraw.New(nil)

	imd.Color = color.NRGBA{0, 0, 0, 255}
	imd.Push(pixel.ZV, pendulum.pendPos1, pendulum.pendPos2)
	imd.Line(3)

	imd.Color = pendulum.color
	imd.Push(pendulum.pendPos1)
	imd.Circle(pendulum.mass1/2, 0)

	imd.Color = pendulum.color
	imd.Push(pendulum.pendPos2)
	imd.Circle(pendulum.mass2/2, 0)

	imd.Draw(win)
}

func (pendulum *sPendulum) a2aCalculation() {
	num1 := 2 * math.Sin(pendulum.theta1-pendulum.theta2)
	num2 := (pendulum.thetaVel1 * pendulum.thetaVel1 * pendulum.radius1 * (pendulum.mass1 + pendulum.mass2))
	num3 := pendulum.grav * (pendulum.mass1 + pendulum.mass2) * math.Cos(pendulum.theta1)
	num4 := pendulum.thetaVel2 * pendulum.thetaVel2 * pendulum.radius2 * pendulum.mass2 * math.Cos(pendulum.theta1-pendulum.theta2)
	den := pendulum.radius2 * (2*pendulum.mass1 + pendulum.mass2 - pendulum.mass2*math.Cos(2*pendulum.theta1-2*pendulum.theta2))

	pendulum.thetaVel2 += (num1 * (num2 + num3 + num4)) / den
}
