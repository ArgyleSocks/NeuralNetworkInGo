package main

import (
  //"fmt"
)

var repetitionSet []bool = make([]bool, len(sampleSet))
var differentSets int = len(sampleSet)
var corresSet [][]int
var corresItem int = 0
var totalSets int = 0

func cleanSamples() {

  for i := 0; i < len(repetitionSet); i++ {
    repetitionSet[i] = false
  }

  for i := 0; i < len(sampleSet); i++ {
    if !repetitionSet[i] {
      for j := i + 1; j < len(sampleSet); j++ {
        if (sampleSet[j][0] == sampleSet[i][0]) {
          repetitionSet[j] = true
          sampleSet[i][1] += sampleSet[j][1]
          differentSets--
        }
      }
    }
  }

  //fmt.Println(differentSets)

  corresSet = make([][]int, differentSets)

  for i := 0; i < len(corresSet); i++ {
    corresSet[i] = make([]int, 2)
  }

  for i := 0; i < len(sampleSet); i++ {
    if !repetitionSet[i] {
      corresSet[corresItem][0] = sampleSet[i][0]
      corresSet[corresItem][1] = sampleSet[i][1]

      for j := i + 1; j < len(sampleSet); j++ {
        if sampleSet[j][0] == sampleSet[i][0] {
          sampleSet[j][0] = corresItem
        }
      }

      sampleSet[i][0] = corresItem

      corresItem++
    }
  }

  for i := 0; i < len(corresSet); i++ {
    totalSets += corresSet[i][1]
  }

  for i := 0; i < len(nodeGraph); i++ {
    for j := 0; j < len(nodeGraph[i]); j++ {
      nodeGraph[i][j].initSums(totalSets)
    }
  }

  /*
  for i := 0; i < len(corresSet); i++ {
    fmt.Print(corresSet[i])
  }
  fmt.Println()

  for i := 0; i < len(sampleSet); i++ {
    fmt.Print(sampleSet[i])
  }
  fmt.Println()

  for i := 0; i < len(sampleSet); i++ {
    fmt.Print(repetitionSet[i])
  }
  fmt.Println()
  */
}
