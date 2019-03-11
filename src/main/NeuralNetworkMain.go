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

type network struct {
  Composition []int
  InputDataSet [][]float64
  ExpectedDataSet [][]float64
  TrainingSet [][2] int
  MinCostRepetition int
  RepValue int
  SampleType int
  RefInputSumType int
  TrainingTask int
  NodeGraph [][]neuron
  TotalSets int
}
// var composition []int
// var words []string  //TODO: Figure out what this is doing in initExpected, as it could probably be diced down to there
// var syllables []int //doesn't need to gloabl TODO: Also figure out what this is as it also doesn't need to bel global but is too weird to touch
//
// var inputDataSet [][]float64
// var expectedDataSet [][]float64 //last entry of expectedDataSet represents 0
// //the indices of inputDataSet and expectedDataSet should correspond to what the input and desired outputs of the NN should be
// var net.TrainingSet [][2]int
// //describes the samples that the network evaluates, 1st index of 2nd dimension is index of case in inputdataSet/expectedData, 2nd index of 2nd dimension is times that case is evaluated
//
// var organizedWords [][]string //can now be
// var organizedLengths []int //can now be optimized to be
//
// //The number of words of any one syllable count a set should contain
//
// //The number of times the established criteria for a "minimum cost" need to be repeated in a row
// var minCostRepetition int = 50
// var repValue = 5 //Knackered
//
// var sampleType int = 1
// var refInputSumType int = 1
// var trainingTask int = 1
//
// //The network graph
// var nodeGraph [][]neuron = make([][]neuron, len(composition))

const UPPER_ASCII_LIM = 122.0
const LOWER_ASCII_LIM = 96.0
const DESIRED_UPPER_ASCII_LIM = 1.0
const DESIRED_LOWER_ASCII_LIM = 0.0

func (net *network) NeuralNetworkExec() {
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

  net.trainNetwork()
  // cleanNetwork()
  //forkManualTest(trainingTask)//TODO remove fork and make to accept array
}

func InitNetworkVar(_composition []int, _inputDataSet [][]float64, _expectedDataSet [][]float64, _trainingSet [][2]int) *network{

  //fmt.Println("Initialization started")
  toR:=new(network)
  toR.Composition = _composition
  toR.InputDataSet = _inputDataSet
  toR.ExpectedDataSet = _expectedDataSet
  toR.TrainingSet = _trainingSet

  // var totalSets := 0 <-- maxim, may I immortalize this extremely em♭arrasing mistake you made
  toR.TotalSets = 0

  for i := 0; i < len(toR.TrainingSet); i++ {
    toR.TotalSets += toR.TrainingSet[i][1]
  }

  for i := 0; i < toR.Composition[compLastRow]; i++ { //need to move this, like this really isn't supposed to be here
    toR.ExpectedDataSet[i] = make([]float64, len(toR.TrainingSet))
  } //difference between expectedDataSet and this: expectedDataSet has the corresponding values to the inputs, this has the values corresponding to the nodegraphs
  toR.NodeGraph = make([][]neuron, len(toR.Composition))
  for i := 0; i < len(toR.Composition); i++ {
    toR.NodeGraph[i] = make([]neuron, toR.Composition[i])
  }

  for i := 0; i < len(toR.Composition); i++ {
    for j := 0; j < toR.Composition[i]; j++ {
      toR.NodeGraph[i][j].initNeuron(i,j)
      toR.NodeGraph[i][j].initSums(len(toR.TrainingSet))//THIS IS ACTUALLY BLATANTLY WRONG! len(toR.TrainingSet) needs to be replaced with THE ACTUAL NUMBER OF SETS. But Josh is lazy right now.
      fmt.Println("LAYER:", toR.NodeGraph[i][j].Layer, "NODE:", toR.NodeGraph[i][j].Node, "weights empty?:", toR.NodeGraph[i][j].Weights == nil)
    }
  }

  // fmt.Println(toR.TrainingSet)

  for i := 0; i < len(toR.InputDataSet); i++ {
    fmt.Println("INPUT", i, ":", toR.InputDataSet[i], "length:", len(toR.InputDataSet[i]))
    fmt.Println("EXPECTED", i, ":", toR.ExpectedDataSet[i], "length:", len(toR.ExpectedDataSet[i]))
  }

  // go drawCostLoop()
  // go drawGraphLoop(&nodeGraph)
  return toR

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

func (net *network) trainNetwork() {

  var firstCost float64 = 0.0
  var lastCost float64 = 0.0
  var generations int = 0
  var endTraining bool = false
  var minCostCheck int = 0

  for train := true; train; train = !endTraining {


    //forkCycle(sampleType)

    bigBoiCycle(net.TrainingSet, net.InputDataSet, net.ExpectedDataSet)

    if generations == 0 {
      net.calcCost(false)
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
