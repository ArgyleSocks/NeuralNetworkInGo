package main

import (
  //"fmt"
  "math"
  // "dict"
)

var compLastRow int = len(composition) - 1
var expected [][]float64 = make([][]float64, composition[compLastRow])
var cost float64
var costDeriv float64

 func setSample(setValue int, set int) {
  //fmt.Println("setValue", setValue)
  //fmt.Println("refSum", len(nodeGraph[0][0].RefInputSum), "set", set)
  switch setValue {
  case 1:
    calcInputNeuron(0, 0, set)
    initExpected(0, 0, set)
  case 2:
    calcInputNeuron(1, 0, set)
    initExpected(1, 0, set)
  case 3:
    calcInputNeuron(0, 1, set)
    initExpected(0, 1, set)
  case 4:
    calcInputNeuron(1, 1, set)
    initExpected(0, 0, set)
  }
}

func initExpected(exp1 float64, exp2 float64, set int) {
	// index:=dict.MapGet(string(word))
	// for i:=0;i<len(expected);i++{
	// 	if i==index{
	// 		expected[i]=1
	// 	} else {
	// 		expected[i]=0
	// 	}
	// }

	expected[0][set] = exp1
	expected[1][set] = exp2
}

func calcCost() {
  cost = 0
  for i := 0; i < len(expected); i++ {
    for j := 0; j < len(expected[0]); j++ {
      cost += math.Pow((nodeGraph[compLastRow][i].RefInputSum[j] - expected[i][j]), 2)
      checkNaN(cost)
    }
  }
}
