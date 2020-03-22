package main

import (
	"fmt"
	"strconv"
)

// Color represents a color
type Color string

// Wheels represents car Wheels
type Wheels string

// Colors contants
const (
	BLUE  Color = "blue"
	GREEN       = "green"
	RED         = "red"
)

// Car is a Car interface
type Car interface {
	Drive() string
	Stop() string
}

// Wheel kind constants
const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

// CarBuilder builds cars
type CarBuilder interface {
	TopSpeed(int) CarBuilder
	Paint(Color) CarBuilder
	Wheels(Wheels) CarBuilder
	Build() Car
}

type carBuilder struct {
	speedOption int
	color       Color
	wheels      Wheels
}

func (cb *carBuilder) Paint(color Color) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) TopSpeed(speed int) CarBuilder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) Wheels(wheels Wheels) CarBuilder {
	cb.wheels = wheels
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.speedOption,
		color:    cb.color,
		wheels:   cb.wheels,
	}
}

// NewCarBuilder creates a CarBuilder
func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

type car struct {
	topSpeed int
	color    Color
	wheels   Wheels
}

func (c *car) Drive() string {
	return "Driving at speed " + strconv.Itoa(c.topSpeed)
}

func (c *car) Stop() string {
	return "Stopping a " + string(c.color) + " car"
}

func main() {
	builder := NewCarBuilder()
	car := builder.TopSpeed(50).Paint(RED).Wheels(SteelWheels).Build()

	fmt.Println(car)
	fmt.Println(car.Drive())
	fmt.Println(car.Stop())
}
