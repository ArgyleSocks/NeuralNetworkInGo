package main

import (
  "math"
)
//nodeGraph: A 2 dimensional array comprised of neurons, presumably the first [] means layer and the second [] is node
//composition: A 1 dimensional array where [] is layer and the values are the number of nodes in [] layer
//composition []int = {300, 900, 100, 50, 15} I think this was it

/* implement stop critera:
All derivative values must be within 0.05 of 0
*/

const trainingRate float64 = 0.2

var costNode, weightLayer, weightNode int = 0, 0, 0
var cycleCount int = 1
var layerDif int = 0
var midNodes []int = make([]int,0)

func backPropPointSelect() {

  for i := 0; i < composition[len(composition) - 1]; i++ {
    for j := 0; j < len(composition) - 1; i++ {
      for k := 0; k < composition[j]; k++ {

        costNode = i
        weightLayer = j
        weightNode = k
        layerDif = len(composition) - (weightLayer + 2)
        midNodes = make([]int, layerDif)

        backPropagation()
      }
    }
  }

  divisor := 1.0

  for i := 0; i < len(composition) - 1; i++ {

    for j := i + 1; j < len(composition); j++ {
      divisor = divisor * float64(composition[i])
    }

    for j := 0; j < composition[i]; i++ {
      for k := 0; k < composition[i + 1]; k++ {
        nodeGraph[i][j].weightsChange[k] = nodeGraph[i][j].weightsChange[k]/divisor
        nodeGraph[i][j].weights[k] -= trainingRate * nodeGraph[i][j].weightsChange[k]
      }
    }
  }

  resetBackPropagation()

}

func backPropagation() {

  if cycleCount <= layerDif {
    for i := 0; i < composition[weightLayer + cycleCount]; i++ {
      midNodes[cycleCount - 1] = i
      cycleCount++
      backPropagation()
    }
  } else {
    weightChange := nodeRefInputSum(weightLayer, weightNode) * sigmoidDerivative(nodeInputSum((weightLayer + 1), midNodes[0]))

    for i := 0; i < layerDif; i++ {
      if i != layerDif - 1 {
        weightChange = weightChange * nodeWeight((weightLayer + i + 1), midNodes[i], midNodes[i + 1]) * sigmoidDerivative(nodeInputSum((weightLayer + i + 2), midNodes[i + 1]))
      } else {
        weightChange = weightChange * nodeWeight((weightLayer + i + 1), midNodes[i], costNode) * sigmoidDerivative(nodeInputSum((weightLayer + i + 2), costNode))
      }
    }

    weightChange = weightChange * 2 * (nodeRefInputSum((len(composition) - 1), costNode) - expected[costNode])

    nodeGraph[weightLayer][weightNode].weightsChange[midNodes[0]] = nodeGraph[weightLayer][weightNode].weightsChange[midNodes[0]] + weightChange
  }
}

func resetBackPropagation() {
  costNode, weightLayer, weightNode = 0, 0, 0
  cycleCount = 1
  layerDif = 0
  midNodes = make([]int,0)
}

func nodeInputSum(layer int, node int) float64{
  return nodeGraph[layer][node].inputSum
}

func nodeRefInputSum(layer int, node int) float64{
  return nodeGraph[layer][node].refInputSum
}

func nodeWeight(layer int, node int, corresNode int) float64{
  return nodeGraph[layer][node].weights[corresNode]
}

func sigmoidDerivative(input float64) float64{
  return 1/(math.Pow((1 + math.Pow(math.E, -input)), 2) * math.Pow(math.E, input))
}
