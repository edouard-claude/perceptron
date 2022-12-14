package perceptron

type Perceptron struct {
	Weights []float64
	Bias    float64
}

func NewPerceptron(n int) *Perceptron {
	p := Perceptron{}
	p.Weights = make([]float64, n)
	p.Bias = 0.0
	return &p
}

func (p *Perceptron) Predict(inputs []float64) int {
	sum := p.Bias
	for i, weight := range p.Weights {
		sum += weight * inputs[i]
	}
	return Step(sum)
}

func (p *Perceptron) Train(inputs []float64, target int) {
	prediction := p.Predict(inputs)
	error := target - prediction
	p.Bias += float64(error)
	for i := range p.Weights {
		p.Weights[i] += float64(error) * inputs[i]
	}
}

func Step(x float64) int {
	if x >= 0 {
		return 1
	}
	return 0
}
