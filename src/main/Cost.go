package main

import (
  "fmt"
  "math"
  // "dict"
)

var compLastRow int = len(composition) - 1
var expected []float64 = make([]float64, composition[compLastRow])
var cost float64
var costDeriv float64

func initExpected() {
	// index:=dict.MapGet(string(word))
	// for i:=0;i<len(expected);i++{
	// 	if i==index{
	// 		expected[i]=1
	// 	} else {
	// 		expected[i]=0
	// 	}
	// }
	expected[0] = 0.5
	expected[1] = 1
}

func calcCost(verbose bool) {
  cost = 0
  for i := 0; i < len(expected); i++ {
  	if verbose {
  	  fmt.Println(nodeGraph[compLastRow][i].RefInputSum)
  	}
    cost += math.Pow((nodeGraph[compLastRow][i].RefInputSum - expected[i]), 2)
    checkNaN(cost)
  }
  if verbose {
    fmt.Println("Cost is ", cost)
  }
}
