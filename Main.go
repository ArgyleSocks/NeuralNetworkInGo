package main

import (
  "fmt"
)

var LPComposition [6] = [...]{2, 3, 3, 3, 3, 2}
var trainingSet1 [4][2]int = [...]{{1, 1}, {2, 1}, {3, 1}, {4, 1}}
var inputDataSet1 [4][2]float64 = [...]{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
var expectedDataSet1 [4][2]float64 = [...]{{0, 0}, {0, 1} {1, 0}, {0, 0}}


func main() {
  fmt.println("Program Started")
  
}
