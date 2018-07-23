package main

import (
  "math"
  "fmt"
)
//nodeGraph: A 2 dimensional array comprised of neurons, presumably the first [] means layer and the second [] is node
//composition: A 1 dimensional array where [] is layer and the values are the number of nodes in [] layer
//composition []int = {300, 900, 100, 50, 15} I think this was it

/* implement stop critera:
All derivative values must be within 0.05 of 0
*/

const trainingRate float64 = 0.02

var endTraining bool = false
var weightLayer, weightNode, weightSelect int = 0, 0, 0
var layerDif int = 0
var threshold float64 = math.Pow(10, -15)

func backPropagation() {

  endTraining = true

  for j := len(composition) - 1; j >= 1; j-- {
    for k := 0; k < composition[j]; k++ {
      for i := 0; i < composition[j - 1]; i++ {
        weightLayer = j
        weightNode = k
        weightSelect = i
        layerDif = len(composition) - (weightLayer + 1)

        nodeGraph[j][k].LocalDeriv = sigmoidDerivative(nodeGraph[j][k].InputSum) * nodeGraph[j-1][i].RefInputSum

        nodeGraph[j][k].TrainRel = true
        calcDerivative(1)
      }
    }
  }

  resetBackPropagation()

}


func calcDerivative(cycleCount int) {
  //fmt.Println("Evaluating backPropagation")
  if cycleCount <= layerDif {
    //fmt.Println("Evaluating if")
    //fmt.Println("layerDif", layerDif, "cycleCount", cycleCount)
    for i := 0; i < composition[len(composition) - cycleCount]; i++ {
      for j := 0; j < composition[len(composition - cycleCount - 1)]; j++ {
        if nodeGraph[len(composition) - cycleCount - 1][j].trainRel {
          nodeGraph[len(composition) - cycleCount][i].LocalDeriv += nodeGraph[len(composition) - cycleCount - 1][j].weights[i] * nodeGraph[len(composition) - cycleCount - 1][j].LocalDeriv * sigmoidDerivative(nodeGraph[len(composition) - cycleCount][i].inputSum)
        }
      }
    }

    backPropagation(cycleCount + 1)

  } else {

    for i := 0; i < compLastRow; i++ {
      costDeriv += 2 * (nodeGraph[compLastRow][i].refInputSum - expected[i]) * nodeGraph[compLastRow][i].LocalDeriv
    }

    nodeGraph[weightLayer - 1][weightSelect].weights[weightNode] -= trainingRate * costDeriv

    if (math.Abs(costDeriv) > threshold) && (!endTraining) {
      endTraining = false
      fmt.Println("Training failed at node", weightLayer, weightSelect + 1, "at weight", weightNode)
    }

  }
}

func resetBackPropagation() {
  weightLayer, weightNode, weightSelect = 0, 0, 0
  layerDif = 0
}

func nodeInputSum(layer int, node int) float64{
  return nodeGraph[layer][node].InputSum
}

func nodeRefInputSum(layer int, node int) float64{
  return nodeGraph[layer][node].RefInputSum
}

func nodeWeight(layer int, node int, corresNode int) float64{
  return nodeGraph[layer][node].Weights[corresNode]
}

func sigmoidDerivative(input float64) float64{
  return 1/(math.Pow((1 + math.Pow(math.E, -input)), 2) * math.Pow(math.E, input))
}