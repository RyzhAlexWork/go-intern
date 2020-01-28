package builder

import "strconv"

type Car interface {
	ShowCar() string
}

// Car implementation.
type car struct {
	wheels      bool
	carcase     bool
	engine      bool
	accelerator bool
	gun         bool
}

// MakeWheels builds a wheels.
func (c *car) MakeWheels() {
	c.wheels = true
}

// MakeCarcase builds a carcase.
func (c *car) MakeCarcase() {
	c.carcase = true
}

// MakeEngine builds a engine.
func (c *car) MakeEngine() {
	c.engine = true
}

// MakeAccelerator builds a accelerator.
func (c *car) MakeAccelerator() {
	c.accelerator = true
}

// MakeGun builds a gun.
func (c *car) MakeGun() {
	c.gun = true
}

// Show returns car.
func (c *car) ShowCar() string {
	return "Wheels:" + strconv.FormatBool(c.wheels) + ". Carcase:" + strconv.FormatBool(c.carcase) +
		". Engine:" + strconv.FormatBool(c.engine) + ". Accelerator:" + strconv.FormatBool(c.accelerator) +
		". Gun:" + strconv.FormatBool(c.gun) + "."
}

func NewCar() Car {
	return &car{}
}
