package main

import (
	//nn "github.com/Forsinge/Yatzy-Neural-Network/nn"
	//yatzy "github.com/Forsinge/Yatzy-Neural-Network/yatzy"
	mat "github.com/Forsinge/Yatzy-Neural-Network/matrix"
)

func main() {

	slice1 := []float64{1, 2, -1, 3, 2, 0, -4, 0, 2}
	slice2 := []float64{3, 4, 2, 0, 1, 0, -2, 0, 1}
	m := mat.NewFrom(3, 3, slice1)
	n := mat.NewFrom(3, 3, slice2)

	m.Print()
	n.Print()

	result := mat.ElemAdd(m, n)
	result.Print()
}
