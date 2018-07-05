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
	expected[7]=1
}

func calcCost() {

  for i := 0; i < len(expected); i++ {
    cost += math.Pow((nodeGraph[compLastRow][i].inputSum - expected[i]), 2)
  }

  fmt.Println("Cost is ", cost)

}
