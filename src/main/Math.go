package main

import (
  "math"
)
//Utility? -->
//Possibly Irrelevant/Needs to be implemented
func nodeInputSum(layer int, node int, set int) float64 {
  return nodeGraph[layer][node].InputSum[set]
}

//Possibly Irrelevant/Needs to be implemented
func nodeRefInputSum(layer int, node int, set int) float64 {
  return nodeGraph[layer][node].RefInputSum[set]
}

//Possibly Irrelevant/Needs to be implemented
func nodeWeight(layer int, node int, corresNode int) float64 {
  return nodeGraph[layer][node].Weights[corresNode]
}

func sigmoid(input float64) float64 {
  return 1 / (1 + (1/(math.Pow(math.E, input))))
}

func sigmoidDerivative(input float64) float64 {
  return (1/(math.Pow((1 + math.Pow(math.E, -input)), 2) * math.Pow(math.E, input)))
}

func trainingRate(slope float64) float64 {
  return (2 * sigmoidDerivative(slope) + 0.02)
}

func ramp(input, lower_lim, upper_lim, desired_lower_lim, desired_upper_lim float64) float64 {
	crange:=upper_lim-lower_lim
	wrange:=desired_upper_lim-desired_lower_lim
	point :=input/crange
	return point*wrange+desired_lower_lim
}
