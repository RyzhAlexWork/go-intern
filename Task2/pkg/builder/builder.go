// Package builder is an example of the Builder Pattern.
package builder

type build struct {
	wheels      bool
	carcase     bool
	engine      bool
	accelerator bool
	gun         bool
}

// Builder provides a builder interface.
type Builder interface {
	MakeWheels() bool
	MakeCarcase() bool
	MakeEngine() bool
	MakeAccelerator() bool
	MakeGun() bool
	BuildCar() Car
}

// MakeWheels builds a wheels.
func (b *build) MakeWheels() bool {
	b.wheels = true
	return true
}

// MakeCarcase builds a carcase.
func (b *build) MakeCarcase() bool {
	b.carcase = true
	return true
}

// MakeEngine builds a engine.
func (b *build) MakeEngine() bool {
	b.engine = true
	return true
}

// MakeAccelerator builds a accelerator.
func (b *build) MakeAccelerator() bool {
	b.accelerator = true
	return true
}

// MakeGun builds a gun.
func (b *build) MakeGun() bool {
	b.gun = true
	return true
}

func (b *build) BuildCar() Car {
	return &car{
		wheels:      b.wheels,
		carcase:     b.carcase,
		engine:      b.engine,
		accelerator: b.accelerator,
		gun:         b.gun,
	}
}

func NewBuild() Builder {
	return &build{}
}
