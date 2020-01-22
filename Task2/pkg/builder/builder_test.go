package builder

import "testing"

var (
	expect = []string{
		"Wheels:true. Carcase:true. Engine:true. Accelerator:true. Gun:false.",
		"Wheels:true. Carcase:true. Engine:true. Accelerator:false. Gun:true.",
		"Wheels:true. Carcase:true. Engine:true. Accelerator:false. Gun:false.",
	}
)

func Test_Builder(t *testing.T) {
	build := NewBuild()
	car := build.BuildCar()
	director := NewDirector(build, car)
	car = director.ConstructSportcar()
	if car.ShowCar() != expect[0] {
		t.Errorf("Wrong answer!")
	}

	build = NewBuild()
	car = build.BuildCar()
	director = NewDirector(build, car)
	car = director.ConstructTank()
	if car.ShowCar() != expect[1] {
		t.Errorf("Wrong answer!")
	}

	build = NewBuild()
	car = build.BuildCar()
	director = NewDirector(build, car)
	car = director.ConstructCar()
	if car.ShowCar() != expect[2] {
		t.Errorf("Wrong answer!")
	}
}
