package perceptron

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestPerceptron(t *testing.T) {
	p := NewPerceptron(2)

	inputs := [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	targets := []int{}
	for _, input := range inputs {
		targets = append(targets, Step(input[0]+input[1]))
	}

	for i := 0; i < 100; i++ {
		for j, input := range inputs {
			p.Train(input, targets[j])
		}
	}

	for _, input := range inputs {
		assert.Equal(t, p.Predict(input), Step(input[0]+input[1]))
	}
}

func TestStep(t *testing.T) {
	assert.Equal(t, Step(0), 1)
	assert.Equal(t, Step(1), 1)
	assert.Equal(t, Step(-1), 0)
}
