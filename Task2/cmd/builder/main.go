package main

import (
	"fmt"
	"github.com/RyzhAlexWork/go-intern/Task2/pkg/builder"
)

func main() {
	build := builder.NewBuild()
	director := builder.NewDirector(build, build.BuildCar())
	car1 := director.ConstructSportcar()
	car2 := director.ConstructTank()

	fmt.Println(car1.ShowCar())
	fmt.Println(car2.ShowCar())
	fmt.Println(car1.ShowCar())
}
