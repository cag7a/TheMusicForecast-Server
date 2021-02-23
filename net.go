package main

import "fmt"

// ActivationFunction is an abstract header for any valid activation function
type ActivationFunction func(float64) float64

// Network is the global model of a Neural Network
type Network struct {
	Name string
}

// InitializeNetwork takes a pointer to a Network and prepares it for training.
// Should always be the first action a Network takes
func InitializeNetwork(net *Network, actType ActivationFunction) {
	if net.Name == "" {
		net.Name = "NewNetwork"
	}

	fmt.Println(net.Name)
}

// LoadWeights takes a pointer to a Network and the name of a weights file to be read in
func LoadWeights(net *Network, fileName string) {

}

// BinaryStep activation function
func BinaryStep(val float64) float64 {
	if val > 0 {
		return 1
	}

	return 0
}
