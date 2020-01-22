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

// Show returns car.
func (c *car) ShowCar() string {
	return "Wheels:" + strconv.FormatBool(c.wheels) + ". Carcase:" + strconv.FormatBool(c.carcase) +
		". Engine:" + strconv.FormatBool(c.engine) + ". Accelerator:" + strconv.FormatBool(c.accelerator) +
		". Gun:" + strconv.FormatBool(c.gun) + "."
}
