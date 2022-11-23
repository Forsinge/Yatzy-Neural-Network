package main

import (
	"fmt"
	"math/rand"
	"time"

	inp "github.com/Forsinge/Yatzy-Neural-Network/input"
	mat "github.com/Forsinge/Yatzy-Neural-Network/matrix"
	nn "github.com/Forsinge/Yatzy-Neural-Network/nn"
	yatzy "github.com/Forsinge/Yatzy-Neural-Network/yatzy"
)

func extractRollOutput(m *mat.Matrix) []int {
	output := []int{}
	for i := 0; i < m.MaxDim(); i += 1 {
		if m.Get(0, i) >= 0.5 {
			output = append(output, i)
		}
	}
	return output
}

func desiredFromInts(input []int) []float64 {
	desired := []float64{0.0, 0.0, 0.0, 0.0, 0.0}
	for i := range input {
		desired[input[i]] = 1.0
	}
	return desired
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	rollNetwork := nn.New([]int{33, 256, 256, 256, 5}, nn.Sigmoid, nn.SigmoidPrime, nn.AdjustRate)
	game := yatzy.NewGame()

	for i := 0; i < 10; i++ {
		inputVec := yatzy.InputVec(game)
		input := mat.NewFrom(len(inputVec), 1, inputVec)

		rollNetwork.ForwardProp(input)
		output := extractRollOutput(rollNetwork.Output())

		game.Print()
		fmt.Println("Output:", output)
		fmt.Print("Enter the correct output: ")

		desired := desiredFromInts(inp.ScanInts())

		fmt.Println("Running backpropagation ...")
		rollNetwork.BackwardProp(mat.NewFrom(len(desired), 1, desired))

		game = yatzy.NewRollState(game)
	}
}
