package main

import (
	"fmt"

	"squirrel/lab/perceptron"
)

func main() {
	p := perceptron.NewPerceptron(2)

	inputs := [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	targets := []int{}
	for _, input := range inputs {
		targets = append(targets, perceptron.Step(input[0]+input[1]))
	}

	for i := 0; i < 100; i++ {
		for j, input := range inputs {
			p.Train(input, targets[j])
		}
	}

	for i, input := range inputs {
		fmt.Printf("Perceptron prediction: %d | Expected: %d | Inputs: %v | Weights: %v | Bias: %v | Error: %v | Iteration: %v | Learning Rate: %v | Activation Function: %v | Loss Function: %v | Cost Function: %v | Gradient Descent: %v | Backpropagation: %v \n", p.Predict(input), perceptron.Step(input[0]+input[1]), input, p.Weights, p.Bias, perceptron.Step(input[0]+input[1])-p.Predict(input), i, 1, "step", "MSE", "MSE", "batch", "backpropagation")
	}
}
