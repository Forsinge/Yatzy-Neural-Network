package nn

import (
	mat "github.com/Forsinge/Yatzy-Neural-Network/matrix"
)

type Network struct {
	layers     []*Layer
	adjustRate func(x float64) float64
	epoch      int
}

// args: size of each layer in order, output activation, derivative of output activation, learning rate adjust
func New(dimensions []int, oa func(x float64) float64, doa func(x float64) float64, ra func(x float64) float64) *Network {
	w := len(dimensions)

	network := &Network{
		layers:     []*Layer{},
		adjustRate: ra,
		epoch:      0,
	}

	for i := 1; i < w-1; i++ {
		layer := NewLayer(dimensions[i-1], dimensions[i], LReLU, LReLUPrime)
		network.layers = append(network.layers, layer)
	}

	layer := NewLayer(dimensions[w-2], dimensions[w-1], oa, doa)
	network.layers = append(network.layers, layer)

	return network
}

func AdjustRate(x float64) float64 {
	return x * 0.01
}

func (nn *Network) Print() {
	for i := range nn.layers {
		nn.layers[i].Print()
	}
}

func (nn *Network) ForwardProp(input *mat.Matrix) *mat.Matrix {
	length := len(nn.layers)
	nn.layers[0].Feed(input)
	for i := 1; i < length; i += 1 {
		nn.layers[i].Feed(nn.layers[i-1].post)
	}

	return nn.layers[length-1].post
}

func (nn *Network) BackwardProp(desired *mat.Matrix) {
	length := len(nn.layers)
	error := mat.ElemSub(desired, nn.layers[length-1].post)
	for i := length - 1; i >= 0; i -= 1 {
		nn.layers[i].Update(error, nn.adjustRate)
		error = nn.layers[i].ErrorProp(error)
	}
}

func (nn *Network) Output() *mat.Matrix {
	return nn.layers[len(nn.layers)-1].post
}
