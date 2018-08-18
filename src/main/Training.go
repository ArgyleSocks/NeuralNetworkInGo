package main

import (
  "math"
  "fmt"
  "strconv"
)
//nodeGraph: A 2 dimensional array comprised of neurons, presumably the first [] means layer and the second [] is node
//composition: A 1 dimensional array where [] is layer and the values are the number of nodes in [] layer
//composition []int = {300, 900, 100, 50, 15} I think this was it

/* implement stop critera:
All derivative values must be within 0.05 of 0
*/

const trainingRate float64 = 0.02

var weightLayer, weightNode, weightSelect, setSelect int = 0, 0, 0, 0
var changeThreshold float64 = 5 * math.Pow(10, -3)
var layerDif int = 0
var stableWeight bool = false

func backPropagation(sets int) {

  stableWeight = true

  for j := len(composition) - 1; j >= 1; j-- {
    for k := 0; k < composition[j]; k++ {
      for i := 0; i < composition[j - 1]; i++ {
        for m := 0; m < sets; m++ {

          weightLayer = j
          weightNode = k
          weightSelect = i
          setSelect = m

          layerDif = len(composition) - (weightLayer + 1)

          nodeGraph[j][k].LocalDeriv = sigmoidDerivative(nodeGraph[j][k].InputSum[setSelect]) * nodeGraph[j-1][i].RefInputSum[setSelect]
          checkNaN(nodeGraph[j][k].LocalDeriv)
          nodeGraph[j][k].TrainRel = true
          calcDerivative(0)
          tempResetBackPropagation()
        }
      }
    }
  }

  for i := 0; i < len(composition) - 1; i++ {
    for j := 0; j < composition[i]; j++ {
      for k := 0; k < composition[i + 1]; k++ {
        fmt.Println("Changing Weight by", trainingRate * (nodeGraph[i][j].WeightsChange[k]/float64(sets))) //don't forget this exists
        //fmt.Println(nodeGraph[i][j].WeightsChange[k])
        nodeGraph[i][j].Weights[k] -= trainingRate * (nodeGraph[i][j].WeightsChange[k]/float64(sets))
      }
    }
  }

  resetBackPropagation()

}

func calcDerivative(cycleCount int) {
  if cycleCount <= layerDif {
    for i := 0; i < composition[weightLayer + cycleCount - 1]; i++ {
      for j := 0; j < composition[weightLayer + cycleCount]; j++ {
        if nodeGraph[weightLayer + cycleCount - 1][j].TrainRel {

          nodeGraph[weightLayer + cycleCount][i].LocalDeriv += nodeGraph[weightLayer + cycleCount - 1][j].Weights[i] * nodeGraph[weightLayer + cycleCount - 1][j].LocalDeriv * sigmoidDerivative(nodeGraph[weightLayer + cycleCount][i].InputSum[setSelect])

          checkNaN(nodeGraph[weightLayer+cycleCount][i].LocalDeriv)

          checkNaN(nodeGraph[weightLayer+cycleCount-1][i].LocalDeriv)
          nodeGraph[weightLayer + cycleCount][i].TrainRel = true
        }
      }
    }

    calcDerivative(cycleCount + 1)

  } else {

    for i := 0; i < composition[compLastRow]; i++ {
      costDeriv += 2 * (nodeGraph[compLastRow][i].RefInputSum[setSelect] - expected[i][setSelect]) * nodeGraph[compLastRow][i].LocalDeriv
    }

    fmt.Println("costDeriv:",costDeriv)
    fmt.Println("WeightsChange before", nodeGraph[weightLayer - 1][weightSelect].WeightsChange[weightNode])
    nodeGraph[weightLayer - 1][weightSelect].WeightsChange[weightNode] += costDeriv
    fmt.Println("WeightsChange after", nodeGraph[weightLayer - 1][weightSelect].WeightsChange[weightNode])

    if (math.Abs(costDeriv) > changeThreshold) && stableWeight {
      stableWeight = false
    }

  }
}

func tempResetBackPropagation() {
  fmt.Println("Temporarily resetting Backpropagation")
  weightLayer, weightNode, weightSelect = 0, 0, 0
  layerDif = 0
  costDeriv = 0.0

  for i := 0; i < len(composition); i++ {
    fmt.Println()
    for j := 0; j < composition[i]; j++ {
      if nodeGraph[i][j].TrainRel {
        fmt.Print(strconv.Itoa(i + 1) + "," + strconv.Itoa(j + 1))
      } else {
         fmt.Print("   ")
      }

      nodeGraph[i][j].TrainRel = false
      nodeGraph[i][j].LocalDeriv = 0

      if !((composition[i] - 1) == j) {
        fmt.Print(" ")
      } else {
        fmt.Println()
      }
    }
  }
}

func resetBackPropagation() {
  for i := 0; i < len(composition) - 1; i++ {
    for j := 0; j < composition[i]; j++ {

      if !(i == (len(composition) - 1)) {
        for k := 0; k < composition[i + 1]; k++ {
          nodeGraph[i][j].WeightsChange[k] = 0
        }
      }
    }
  }
}

func nodeInputSum(layer int, node int, set int) float64{
  return nodeGraph[layer][node].InputSum[set]
}

func nodeRefInputSum(layer int, node int, set int) float64{
  return nodeGraph[layer][node].RefInputSum[set]
}

func nodeWeight(layer int, node int, corresNode int) float64{
  return nodeGraph[layer][node].Weights[corresNode]
}

func sigmoidDerivative(input float64) float64{
  return 1/(math.Pow((1 + math.Pow(math.E, -input)), 2) * math.Pow(math.E, input))
}
