package main

import (
  "fmt"
  // "math"
  "dict"
  "time"
  // "reflect"
  "math/rand"
)

var compLastRow int = len(composition) - 1 //Turns out this helps in context
var expected [][]float64 = make([][]float64, composition[compLastRow])
var cost float64 //doesn't need to global

func initExpected(expectedResult []float64, setIndex int) {//supposed to set expected, but was converted to do the job setSample really does, but since setSample does it, it is obsolete. We still need to set expected.
  for i, e := range expectedResult {
    expected[i][setIndex] = 0
    expected[i][setIndex] = e
  }
  //expected=append(expected,expectedSampleResult)  OBSOLETE
}
