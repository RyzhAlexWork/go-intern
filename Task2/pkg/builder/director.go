package builder

type builder interface {
	MakeWheels() bool
	MakeCarcase() bool
	MakeEngine() bool
	MakeAccelerator() bool
	MakeGun() bool
	BuildCar() car
}

type Director interface {
	ConstructCar() Car
	ConstructSportcar() Car
	ConstructTank() Car
}

// Director implements a manager
type director struct {
	builder Builder
	car     Car
}

// Construct car.
func (d *director) ConstructCar() Car {
	d.builder.MakeWheels()
	d.builder.MakeCarcase()
	d.builder.MakeEngine()
	return d.builder.BuildCar()
}

// Construct sport-car.
func (d *director) ConstructSportcar() Car {
	d.builder.MakeWheels()
	d.builder.MakeCarcase()
	d.builder.MakeEngine()
	d.builder.MakeAccelerator()
	return d.builder.BuildCar()
}

// Construct tank.
func (d *director) ConstructTank() Car {
	d.builder.MakeWheels()
	d.builder.MakeCarcase()
	d.builder.MakeEngine()
	d.builder.MakeGun()
	return d.builder.BuildCar()
}

func NewDirector(inputBuilder Builder, inputCar Car) Director {
	return &director{
		builder: inputBuilder,
		car:     inputCar,
	}
}
