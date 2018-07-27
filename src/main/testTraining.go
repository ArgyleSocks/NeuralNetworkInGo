package main

import (
  "math"
  "fmt"
  //"strconv"
)
//nodeGraph: A 2 dimensional array comprised of neurons, presumably the first [] means layer and the second [] is node
//composition: A 1 dimensional array where [] is layer and the values are the number of nodes in [] layer
//composition []int = {300, 900, 100, 50, 15} I think this was it

/* implement stop critera:
All derivative values must be within 0.05 of 0
*/

const trainingRate float64 = 0.02
var threshold float64 = math.Pow(10, -15)

var endTraining bool = false
var weightLayer, weightNode, weightSelect int = 0, 0, 0
var layerDif int = 0

func backPropagation() {

  //fmt.Println("Swurlk")
  endTraining = true

  for j := len(composition) - 1; j >= 1; j-- {
    for k := 0; k < composition[j]; k++ {
      for i := 0; i < composition[j - 1]; i++ {
        weightLayer = j
        weightNode = k
        weightSelect = i
        layerDif = len(composition) - (weightLayer + 1)

        nodeGraph[j][k].LocalDeriv = sigmoidDerivative(nodeGraph[j][k].InputSum) * nodeGraph[j-1][i].RefInputSum
        fmt.Println("checking nodeGraph[", j, "][", k, "].LocalDeriv (nodeGraph[j][k].LocalDeriv)")
        checkNaN(nodeGraph[j][k].LocalDeriv)
        nodeGraph[j][k].TrainRel = true
        calcDerivative(0)
        //fmt.Println("Doink Grips")
        resetBackPropagation()
      }
    }
  }
}

func calcDerivative(cycleCount int) {
  if cycleCount <= layerDif {
    //fmt.Println("weightLayer:", weightLayer, "layerDif:", layerDif, "Cyclecount:", cycleCount)
    for i := 0; i < composition[weightLayer + cycleCount - 1]; i++ {
      for j := 0; j < composition[weightLayer + cycleCount]; j++ {
        if nodeGraph[weightLayer + cycleCount - 1][j].TrainRel {
          nodeGraph[weightLayer + cycleCount][i].LocalDeriv += nodeGraph[weightLayer + cycleCount - 1][j].Weights[i] * nodeGraph[weightLayer + cycleCount - 1][j].LocalDeriv * sigmoidDerivative(nodeGraph[weightLayer + cycleCount][i].InputSum)
          fmt.Println("Checking LocalDeriv at nodeGraph[", (weightLayer + cycleCount), "][", i, "].LocalDeriv (nodeGraph[weightLayer + cycleCount][i].LocalDeriv)")
          checkNaN(nodeGraph[weightLayer+cycleCount][i].LocalDeriv)
          fmt.Println("Checking LocalDeriv at nodeGraph[", (weightLayer + cycleCount - 1), "][", i, "].LocalDeriv (nodeGraph[weightLayer + cycleCount - 1][i].LocalDeriv)")
          checkNaN(nodeGraph[weightLayer+cycleCount-1][i].LocalDeriv)
          nodeGraph[weightLayer + cycleCount][i].TrainRel = true
        }
      }
    }

    calcDerivative(cycleCount + 1)

  } else {

    for i := 0; i < composition[compLastRow]; i++ {
      costDeriv += 2 * (nodeGraph[compLastRow][i].RefInputSum - expected[i]) * nodeGraph[compLastRow][i].LocalDeriv
    }

    nodeGraph[weightLayer - 1][weightSelect].Weights[weightNode] -= trainingRate * costDeriv

    if (math.Abs(costDeriv) > threshold) && (!endTraining) {
      endTraining = false
      fmt.Println("Training failed at node", weightLayer, weightSelect + 1, "at weight", weightNode)
    }

  }
}

func resetBackPropagation() {
  weightLayer, weightNode, weightSelect = 0, 0, 0
  layerDif = 0
  costDeriv = 0.0

  /*for i := 0; i < len(composition); i++ {
    fmt.Println()
    for j := 0; j < composition[i]; j++ {
      if nodeGraph[i][j].TrainRel {
        fmt.Print(strconv.Itoa(i + 1) + "," + strconv.Itoa(j + 1))
      } else {
        fmt.Print("   ")
      }

      nodeGraph[i][j].TrainRel = false
      nodeGraph[i][j].LocalDeriv = 0

      fmt.Println("I at line 95:", i, "J at line 95:", j)

      if !((composition[i] - 1) == j) {
        fmt.Print(" ")
      }
    }
  }*/
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
