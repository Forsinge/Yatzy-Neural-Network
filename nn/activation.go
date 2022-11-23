package nn

import (
	"math"
)

func LReLU(x float64) float64 {
	if x >= 0 {
		return x
	} else {
		return 0.1 * x
	}
}

func LReLUPrime(x float64) float64 {
	if x >= 0 {
		return 1
	} else {
		return 0.1
	}
}

func Sigmoid(x float64) float64 {
	exp := math.Pow(math.E, x)
	return exp / (exp + 1)
}

func SigmoidPrime(y float64) float64 {
	return Sigmoid(y) * (1 - Sigmoid(y))
}
