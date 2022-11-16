package nn

import (
	mat "github.com/Forsinge/Yatzy-Neural-Network/matrix"
)

type Layer struct {
	input      *mat.Matrix
	weights    *mat.Matrix // output rows, input cols
	biases     *mat.Matrix
	pre        *mat.Matrix
	post       *mat.Matrix
	activation func(x float64) float64
	derivative func(x float64) float64
}

func New(inputSize, outputSize int, activation func(x float64) float64, derivative func(x float64) float64) *Layer {
	return &Layer{
		weights:    mat.NewRand(outputSize, inputSize),
		biases:     mat.NewRand(outputSize, 1),
		activation: activation,
		derivative: derivative,
	}
}

func (l *Layer) Feed(input *mat.Matrix) {
	l.input = input
	mul := mat.Mul(l.weights, input)
	l.pre = mat.ElemAdd(mul, l.biases)
	l.post = mat.Copy(l.pre)
	l.post.Apply(l.activation)
}

func (l *Layer) Update(error *mat.Matrix, adjustRate func(x float64) float64) {
	l.input.Transpose()
	dZ := mat.ElemMul(error, mat.Apply(l.pre, l.derivative))
	dZ.Apply(adjustRate)
	dW := mat.Mul(dZ, l.input)
	l.input.Transpose()

	l.weights.Sub(dW)
	l.biases.Sub(dZ)
}

func (l *Layer) ErrorProp(error *mat.Matrix) *mat.Matrix {
	l.weights.Transpose()
	ret := mat.Mul(l.weights, error)
	l.weights.Transpose()
	return ret
}
