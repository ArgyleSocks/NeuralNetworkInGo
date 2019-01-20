package main

import (
  "math"
  //"fmt"
  //"strconv"
)
//nodeGraph: A 2 dimensional array comprised of neurons, presumably the first [] means layer and the second [] is node
//composition: A 1 dimensional array where [] is layer and the values are the number of nodes in [] layer
//composition []int = {300, 900, 100, 50, 15} I think this was it

/* implement stop criteria:
All derivative values must be within 0.05 of 0
*/

//trainingRate was 0.02

var weightLayer  int = 0//BIG TODO: REVAMP THIS TO BE LOCAL
var weightNode   int = 0//And this
var weightSelect int = 0//And this
var setSelect    int = 0//And this

var changeThreshold float64 = 5 * math.Pow(10, -3)
var layerDif int = 0 //BIG TODO: must make local

//HMMMMMM
var costDeriv float64 //BIG TODO: must make local
//we have a local costDeriv, it's LocalDeriv. So do we need this? if it's referenced anywhere it'll be a pain with paralellization

var stableWeight bool = false

func backPropagation(sets int) {

  stableWeight = true

  for i := len(composition) - 1; i >= 1; i-- { //BIG TODO: Move to an external function and have it pass values to backPropagation itself
    for j := 0; j < composition[i]; j++ { //TODO
      for k := 0; k < composition[i - 1]; k++ { //TODO
        for m := 0; m < sets; m++ { //TODO

          weightLayer = i
          weightNode = j
          weightSelect = k
          setSelect = m

          layerDif = len(composition) - (weightLayer + 1)

          //what is LocalDeriv? What is RefInputSum? Explain it ALLLLLLLLLLLLLL
          nodeGraph[i][j].LocalDeriv = forkDerivative(refInputSumType, nodeGraph[i][j].InputSum[setSelect]) * nodeGraph[i-1][k].RefInputSum[m]
          checkNaN(nodeGraph[i][j].LocalDeriv)
          nodeGraph[i][j].TrainRel = true
          calcDerivative(0)
          tempResetBackPropagation()
        }
        //nodeGraph[i - 1][k].Weights[j] -= trainingRate(nodeGraph[i][k].WeightsChange[j]/float64(sets)) * (nodeGraph[i][k].WeightsChange[j]/float64(sets))

      }
    }
  }
  //part 2 what does part 2 do?
  for i := 0; i < len(composition) - 1; i++ {//per layer
    for j := 0; j < composition[i]; j++ {//per node
      for k := 0; k < composition[i + 1]; k++ {//per node next layer, making permutation with previous for loop
        //fmt.Println("Changing Weight by", trainingRate * (nodeGraph[i][j].WeightsChange[k]/float64(sets))) //don't forget this exists
        //fmt.Println(nodeGraph[i][j].WeightsChange[k])
        nodeGraph[i][j].Weights[k] -= trainingRate(nodeGraph[i][j].WeightsChange[k]/float64(sets)) * (nodeGraph[i][j].WeightsChange[k]/float64(sets))
      }
    }
  } //this

  //what does this do?
  resetBackPropagation()

}

func calcDerivative(cycleCount int) {
  if cycleCount <= layerDif {
    for i := 0; i < composition[weightLayer + cycleCount - 1]; i++ {
      for j := 0; j < composition[weightLayer + cycleCount]; j++ {
        if nodeGraph[weightLayer + cycleCount - 1][j].TrainRel {

          nodeGraph[weightLayer + cycleCount][i].LocalDeriv += nodeGraph[weightLayer + cycleCount - 1][j].Weights[i] * nodeGraph[weightLayer + cycleCount - 1][j].LocalDeriv * forkDerivative(refInputSumType, nodeGraph[weightLayer + cycleCount][i].InputSum[setSelect])

          //checkNaN(nodeGraph[weightLayer+cycleCount][i].LocalDeriv)

          //checkNaN(nodeGraph[weightLayer+cycleCount-1][i].LocalDeriv)
          nodeGraph[weightLayer + cycleCount][i].TrainRel = true
        }
      }
    }

    calcDerivative(cycleCount + 1)

  } else {

    for i := 0; i < composition[compLastRow]; i++ {
      costDeriv += 2 * (nodeGraph[compLastRow][i].RefInputSum[setSelect] - expected[i][setSelect]) * nodeGraph[compLastRow][i].LocalDeriv
    }

    //fmt.Println("WeightsChange before", nodeGraph[weightLayer - 1][weightSelect].WeightsChange[weightNode])
    nodeGraph[weightLayer - 1][weightSelect].WeightsChange[weightNode] += costDeriv
    //fmt.Println("WeightsChange after", nodeGraph[weightLayer - 1][weightSelect].WeightsChange[weightNode])

    if (math.Abs(costDeriv) > changeThreshold) && stableWeight {
      stableWeight = false
    }
  }
}

func tempResetBackPropagation() {
  //fmt.Println("Temporarily resetting Backpropagation")
  weightLayer, weightNode, weightSelect, setSelect, layerDif = 0, 0, 0, 0, 0
  // layerDif = 0
  costDeriv = 0.0

  //fmt.Println("SAMPLE", setSelect)

  for i := 0; i < len(composition); i++ {
    //fmt.Println()
    for j := 0; j < composition[i]; j++ {
      /*if nodeGraph[i][j].TrainRel {
        fmt.Print(strconv.Itoa(i + 1) + "," + strconv.Itoa(j + 1))
      } else {
         fmt.Print("   ")
      }*/

      nodeGraph[i][j].TrainRel = false
      nodeGraph[i][j].LocalDeriv = 0

      /*if !((composition[i] - 1) == j) {
        fmt.Print(" ")
      } else {
        fmt.Println()
      }*/
    }
  }
}
//mention the difference between tempReset and reset
func resetBackPropagation() {
  /*for i := 0; i < len(composition) - 1; i++ {
    for j := 0; j < composition[i]; j++ {
      if !(i == compLastRow) {
        for k := 0; k < composition[i + 1]; k++ {
          nodeGraph[i][j].WeightsChange[k] = 0
        }
      }
    }
  }*/

  for i := 0; i < len(composition) - 1; i++ {
    for j := 0; j < composition[i]; j++ {
      if !(i == compLastRow) {
        for k := 0; k < composition[i + 1]; k++ {
          nodeGraph[i][j].WeightsChange[k] = 0
        }
      }
    }
  }
}
