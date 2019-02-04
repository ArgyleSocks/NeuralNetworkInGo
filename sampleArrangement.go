package main

import (
  "fmt"
)

var totalSets int = 0
var corresSet [][]int

func twoDiCleanup() {
  // I don't know how to comment this yet, I need to consult ArgyleSocks
  var repetitionSet []bool = make([]bool, len(trainingSet))
  var differentSets int = 0
  var corresItem int = 0

  for i := 0; i < len(repetitionSet); i++ {
    repetitionSet[i] = false
  }

  for i := 0; i < len(trainingSet); i++ {
    if !repetitionSet[i] {
      differentSets++
      for j := i + 1; j < len(trainingSet); j++ {
        if (trainingSet[j][0] == trainingSet[i][0]) {
          repetitionSet[j] = true
          trainingSet[i][1] += trainingSet[j][1]
        }
      }
    }
  }
  //----------Might need to help me out here-----------
/*
  differentSets = len(words/*this can also be the length of the map or some gen. value for number of words*//*)*/

  fmt.Println("DifferentSets:", differentSets)

  corresSet = make([][]int, differentSets)

  for i := 0; i < len(corresSet); i++ {
    corresSet[i] = make([]int, 2)
  }

  for i := 0; i < len(trainingSet); i++ {
    if !repetitionSet[i] {
      fmt.Println("len(corresSet):", len(corresSet), "corresItem:", corresItem, "i:", i, "len(trainingSet):", len(trainingSet), "len(trainingSet[i]):", len(trainingSet[i]), "trainingSet[i]:", trainingSet[i])
      corresSet[corresItem][0] = trainingSet[i][0]
      corresSet[corresItem][1] = trainingSet[i][1]

      for j := i + 1; j < len(trainingSet); j++ {
        if trainingSet[j][0] == trainingSet[i][0] {
          trainingSet[j][0] = corresItem
        }
      }

      trainingSet[i][0] = corresItem

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

  for i := 0; i < len(trainingSet); i++ {
    fmt.Print(trainingSet[i])
  }
  fmt.Println()

  for i := 0; i < len(trainingSet); i++ {
    fmt.Print(repetitionSet[i])
  }
  fmt.Println() */
}

func uniformCasesCleanup() {

  maxSyllable := 0
  varietySyllable := 0
  var repeatCheck []bool = make([]bool, len(syllables))

  //This figures out how many different syllable counts there are among the words
  for i := 0; i < len(syllables); i++ {
    if !repeatCheck[i] {

      varietySyllable++

      for k := i + 1; k < len(syllables); k++ {
        if syllables[k] == syllables[i] {
          repeatCheck[k] = true
        }
      }

      if syllables[i] > maxSyllable {
        maxSyllable = syllables[i]
      }
    }
  }

  organizedLengths = make([]int, varietySyllable)
  organizedWords = make([][]string, varietySyllable)

  typeSyllables := 0
  columnTick := 0
  rowTick := 0


  for i := 0; i < len(repeatCheck); i++ {
    if !repeatCheck[i] {
      organizedLengths[columnTick] = syllables[i]

      typeSyllables = 1
      for k := i + 1; k < len(syllables); k++ {
        if syllables[k] == syllables[i] {
          typeSyllables++
        }
      }

      organizedWords[columnTick] = make([]string, typeSyllables)

      for k := i; k < len(syllables); k++ {
        if syllables[k] == syllables[i] {
          organizedWords[columnTick][rowTick] = words[k]
          rowTick++
        }
      }

      columnTick++
      rowTick = 0
      typeSyllables = 0
    }
  }

  totalSets = len(organizedLengths) * repValue

  for i := 0; i < len(nodeGraph); i++ {
    for j := 0; j < len(nodeGraph[i]); j++ {
      nodeGraph[i][j].initSums(totalSets)
    }
  }
}

func bigBoiCycle(trainingSet [][2]int, inputDataSet, wanted [][]float64) {
/*from the "developer" side, by which I mean people who do not see this library's internals,
one may call this with two 2d arrays -- every index in such array corresponds to a sample,
and the second dimension allows for multiple inputs/outputs.
Make sure that all samples are formatted to the correct length.*/
  setCounter := 0
  for i := 0; i < len(trainingSet); i++ {
    for j := 0; j < trainingSet[i][1]; j++ {
      fmt.Println(trainingSet)
      calcInputNeuron(inputDataSet[trainingSet[i][0]],setCounter)
      initExpected(wanted[trainingSet[i][0]],setCounter)
      evaluateNetwork(setCounter)//whoo, nearly forgot that one
      setCounter++
    }
  }
}

func twoDiCycle() {
  var setCounter = 0
  fmt.Println(corresSet)
  for i := 0; i < len(corresSet); i++ {
    for j := 0; j < corresSet[i][1]; j++ {
      //fmt.Println("j", j)
      forkTrainingTask(trainingTask, corresSet[i][0], setCounter)
      evaluateNetwork(setCounter)
      setCounter++
    }
  }
}

func uniformCasesCycle() {
  var setCounter = 0

  for i := 0; i < len(organizedWords); i++ {
    for k := 0; k < repValue; k++ {
      forkTrainingTask(trainingTask, i, setCounter)
      evaluateNetwork(setCounter)
      setCounter++
    }
  }
}
