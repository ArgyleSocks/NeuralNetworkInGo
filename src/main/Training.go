package main

import (
  "math"
  //"fmt"
)
//nodeGraph: A 2 dimensional array comprised of neurons, presumably the first [] means layer and the second [] is node
//composition: A 1 dimensional array where [] is layer and the values are the number of nodes in [] layer
//composition []int = {300, 900, 100, 50, 15} I think this was it

/* implement stop critera:
All derivative values must be within 0.05 of 0
*/

const trainingRate float64 = 0.2

var endTraining bool = false
var weightLayer, weightNode int = 0, 0
var layerDif int = 0
var midNodes []int = make([]int, 0)
var divisor = 1.0

func backPropPointSelect() {
  for j := 0; j < len(composition) - 1; j++ {
    for k := 0; k < composition[j]; k++ {

      weightLayer = j
      weightNode = k
      layerDif = len(composition) - (weightLayer + 1)
      midNodes = make([]int, layerDif)
      //fmt.Println("SOME COOL STUFF",weightLayer,weightNode,"\n",composition[4],"OR NOT")
      backPropagation(1)

    }
  }

  for i := 0; i < len(composition) - 1; i++ {

    for j := i + 1; j < len(composition); j++ {
      divisor = divisor * float64(composition[i])
    }

    for j := 0; j < composition[i]; i++ {
      for k := 0; k < composition[i + 1]; k++ {
        nodeGraph[i][j].weightsChange[k] = nodeGraph[i][j].weightsChange[k]/divisor
        nodeGraph[i][j].weights[k] -= trainingRate * nodeGraph[i][j].weightsChange[k]

        endTraining = true
        if ((-nodeGraph[i][j].weightsChange[k] > 0.01) || (-nodeGraph[i][j].weightsChange[k] < -0.01)) && (!endTraining) {
          endTraining = false
          //if all of the weights become finely tuned enough that the changes required are within +-0.01 of 0 (even less than that, actually), the program stops training
          //can still plateau, is still an issue that needs to be resolved https://www.desmos.com/calculator/0hhji76otn
        }
      }
    }
  }

  resetBackPropagation()

}

func backPropagation(cycleCount int) {
  //fmt.Println("Evaluating backPropagation")
  if cycleCount <= layerDif {
    //fmt.Println("Evaluating if")
    //fmt.Println("layerDif", layerDif, "cycleCount", cycleCount)
    for i := 0; i < composition[weightLayer + cycleCount]; i++ {
      midNodes[cycleCount - 1] = i
      //fmt.Println("cycles",cycleCount)
      backPropagation(cycleCount + 1)
    }
  } else {

    //fmt.Println("Evaluating else")

    weightChange := nodeRefInputSum(weightLayer, weightNode) * sigmoidDerivative(nodeInputSum((weightLayer + 1), midNodes[0]))

    //fmt.Println("Argyle")

    for i := 0; i < layerDif -1; i++ {
      //fmt.Println("Sock", i)
      weightChange = weightChange * nodeWeight((weightLayer + i + 1), midNodes[i], midNodes[i + 1]) * sigmoidDerivative(nodeInputSum((weightLayer + i + 2), midNodes[i + 1]))
    }

    //fmt.Println("X gon give it to ya")

    weightChange = weightChange * 2 * (nodeRefInputSum((len(composition) - 1), midNodes[layerDif - 1]) - expected[midNodes[layerDif - 1]])

    nodeGraph[weightLayer][weightNode].weightsChange[midNodes[0]] = nodeGraph[weightLayer][weightNode].weightsChange[midNodes[0]] + weightChange

  }
}

func resetBackPropagation() {
  weightLayer, weightNode = 0, 0
  layerDif = 0
  midNodes = make([]int,0)
  divisor = 1.0
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
