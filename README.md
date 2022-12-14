# Golang Perceptron implementation

### Usage

```sh
go run cmd/sample/main.go
```

```sh
Perceptron prediction: 1 | Expected: 1 | Inputs: [0 0] | Weights: [0 0] | Bias: 0 | Error: 0 | Iteration: 0 | Learning Rate: 1 | Activation Function: step | Loss Function: MSE | Cost Function: MSE | Gradient Descent: batch | Backpropagation: backpropagation 
Perceptron prediction: 1 | Expected: 1 | Inputs: [0 1] | Weights: [0 0] | Bias: 0 | Error: 0 | Iteration: 1 | Learning Rate: 1 | Activation Function: step | Loss Function: MSE | Cost Function: MSE | Gradient Descent: batch | Backpropagation: backpropagation 
Perceptron prediction: 1 | Expected: 1 | Inputs: [1 0] | Weights: [0 0] | Bias: 0 | Error: 0 | Iteration: 2 | Learning Rate: 1 | Activation Function: step | Loss Function: MSE | Cost Function: MSE | Gradient Descent: batch | Backpropagation: backpropagation 
Perceptron prediction: 1 | Expected: 1 | Inputs: [1 1] | Weights: [0 0] | Bias: 0 | Error: 0 | Iteration: 3 | Learning Rate: 1 | Activation Function: step | Loss Function: MSE | Cost Function: MSE | Gradient Descent: batch | Backpropagation: backpropagation 
```

### References

- [Perceptron](https://en.wikipedia.org/wiki/Perceptron)
- [Perceptron Learning Algorithm](https://en.wikipedia.org/wiki/Perceptron#Learning_algorithm)
- [Perceptron Learning Rule](https://en.wikipedia.org/wiki/Perceptron#Learning_rule)
