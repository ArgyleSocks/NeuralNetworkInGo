package main

import (
  //"os"
  "fmt"
  // "time"
  // "dict" 🎸
  //"bufio"
  // "runtime"
  // "strconv"
  // "strings"
  // "math/rand"
)

//TODO: Make the layer size auto adjust to the length of the longest word
//TODO: Threading
//TODO: Cleanup and commenting
//TODO: Investigate the error where differences in layer size between layers cause an index out of range

var composition []int
var words []string  //TODO: Figure out what this is doing in initExpected, as it could probably be diced down to there
var syllables []int //doesn't need to gloabl TODO: Also figure out what this is as it also doesn't need to bel global but is too weird to touch

var inputDataSet [][]float64
var expectedDataSet [][]float64
//the indices of inputDataSet and expectedDataSet should correspond to what the input and desired outputs of the NN should be
var trainingSet [][2]int
//describes the samples that the network evaluates, 1st index of 2nd dimension is index of case in inputdataSet/expectedData, 2nd index of 2nd dimension is times that case is evaluated

var organizedWords [][]string //can now be
var organizedLengths []int //can now be optimized to be

//The number of words of any one syllable count a set should contain

//The number of times the established criteria for a "minimum cost" need to be repeated in a row
var minCostRepetition int = 50
var repValue = 5 //Knackered

var sampleType int = 1
var refInputSumType int = 1
var trainingTask int = 1

//The network graph
var nodeGraph [][]neuron = make([][]neuron, len(composition))

const UPPER_ASCII_LIM = 122.0
const LOWER_ASCII_LIM = 45.0
const DESIRED_UPPER_INPUT_LIM = 1.0
const DESIRED_LOWER_INPUT_LIM = 0.0

func NeuralNetworkExec() {
  // runtime.GOMAXPROCS(1024)
  // dict.Initi("../../dat/syllables")
  //shows where the syllables file is
  //TODO: variadic such that me and maxim don't have to swap it back and forth when either want to run it.
  // dict.ToMap()

  //Maps are faster than array iteration, believe it or not.

  //initi()
  //Initialization
  //prepares graph and samples

  //need to fit in bigBoiCycle...
  //

  trainNetwork()
  // cleanNetwork()
  forkManualTest(trainingTask)//TODO call from Main.go
}

func InitNetworkVar(_composition []int, _inputDataSet [][]float64, _expectedDataSet [][]float64, _trainingSet [][2]int) {

  fmt.Println("Initialization started")

  composition = _composition
  inputDataSet = _inputDataSet
  expectedDataSet = _expectedDataSet
  trainingSet = _trainingSet

  // var totalSets := 0 <-- maxim, may I immortalize this extremely em♭arrasing mistake you made
  totalSets := 0

  for i := 0; i < len(trainingSet); i++ {
    totalSets += trainingSet[i][1]
  }

  for i := 0; i < composition[compLastRow]; i++ { //need to move this, like this really isn't supposed to be here
    expected[i] = make([]float64, len(trainingSet))
  } //difference between expectedDataSet and this: expectedDataSet has the corresponding values to the inputs, this has the values corresponding to the nodegraphs
  nodeGraph=make([][]neuron,len(composition))
  for i := 0; i < len(composition); i++ {
    nodeGraph[i] = make([]neuron, composition[i])
    fmt.Println(nodeGraph[i])
  }

  for i := 0; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].initNeuron(i,j)
      nodeGraph[i][j].initSums(len(trainingSet))//THIS IS ACTUALLY BLATANTLY WRONG! len(trainingSet) needs to be replaced with THE ACTUAL NUMBER OF SETS. But Josh is lazy right now.
    }
  }

  go drawCostLoop()
  go drawGraphLoop(&nodeGraph)

}

/*func initi() {

  //TODO words = dict.SetOfKeys()
  //TODO letterCountSampleVariety := len(words)

  for i := 0; i < len(composition); i++ {
    nodeGraph[i] = make([]neuron, composition[i])
  }

  for i := 0; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].initNeuron(i,j)
    }
  }

  //TODO wordLengths := make([]int, letterCountSampleVariety)
  //TODO for i := 0; i < letterCountSampleVariety; i++ {
    //TODO syllables[i] = dict.MapGet(words[i])
    //TODO this is probably the more relevant one wordLengths[i] = len(words[i])
  }

  forkCleanup(sampleType) //TODO: Actually annihilate all forks

}*/

func trainNetwork() {

  var firstCost float64 = 0.0
  var lastCost float64 = 0.0
  var generations int = 0
  var endTraining bool = false
  var minCostCheck int = 0

  for train := true; train; train = !endTraining {


    //forkCycle(sampleType)

    bigBoiCycle(trainingSet,inputDataSet, expectedDataSet)

    if generations == 0 {
      calcCost(false)
      fmt.Println("cost:",cost)
      firstCost = cost
    }

    backPropagation(totalSets)
    calcCost(false)
    if (cost == lastCost) || stableWeight {
      minCostCheck++
      if minCostCheck >= minCostRepetition {
        endTraining = true
      }
    } else {
      minCostCheck = 0
    }

    fmt.Println("COST:", cost, "CHANGE:", (cost - lastCost), "GENERATION:", generations)

    lastCost = cost
    generations++

    if (cost - lastCost) > 0 || cost < 0.5 {
      fmt.Println(cost-lastCost, cost)
    }

  }

  fmt.Println("gen", generations)

  calcCost(false)

  fmt.Println("First cost:", firstCost, "\b, Last cost:", lastCost)
  fmt.Println("Change in cost:", (lastCost - firstCost) )
  //cleanup
  endTraining = false
  generations = 0
  minCostCheck = 0
}

func evaluateNetwork(graph int) {
  for i := 1; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].calcInputSum(graph)
    }
  }
}

func cleanNetwork() {
  for i := 0; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      for m := 0; m < totalSets; m++ {
        nodeGraph[i][j].RefInputSum[m] = 0
        nodeGraph[i][j].InputSum[m] = 0
        nodeGraph[i][j].OutputSum[m] = 0
      }
    }
  }
}
