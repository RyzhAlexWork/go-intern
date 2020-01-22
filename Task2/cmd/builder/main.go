package main

import (
	"fmt"

	"github.com/RyzhAlexWork/go-intern/Task2/pkg/builder"
)

func main() {
	build := builder.NewBuild()
	car := build.BuildCar()
	director := builder.NewDirector(build, car)
	car = director.ConstructSportcar()

	fmt.Println(car.ShowCar())
}
