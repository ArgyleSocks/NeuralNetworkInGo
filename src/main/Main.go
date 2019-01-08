package main

import (
  "os"
  "fmt"
  // "time"
  "dict"
  "bufio"
  // "runtime"
  // "strconv"
  // "strings"
  // "math/rand"
)

//TODO: Make the layer size auto adjust to the length of the longest word
//TODO: Threading
//TODO: Cleanup and commenting
//TODO: Investigate the error where differences in layer size between layers cause an index out of range

var composition[6]int = [...]int{30, 30, 30, 30, 30, 30}
var sampleSet [5][2]int = [...][2]int{{1, 20}, {2, 20}, {3, 20}, {4, 20}, {5, 20}} //if you use this make sure to use clean samples 1
var words []string  //TODO: Figure out what this is doing in initExpected, as it could probably be diced down to there
var syllables []int //doesn't need to gloabl TODO: Also figure out what this is as it also doesn't need to bel global but is too weird to touch

var organizedWords [][]string
var organizedSyllables []int

//The number of words of any one syllable count a set should contain

//The number of times the established criteria for a "minimum cost" need to be repeated in a row
var minCostRepetition int = 50
var repValue = 5

var inputType int = 2
var refInputSumType int = 1

//The network graph
var nodeGraph [][]neuron = make([][]neuron, len(composition))

const UPPER_LIM = 122.0
const LOWER_LIM = 45.0
const DESIRED_UPPER_LIM = 1.0
const DESIRED_LOWER_LIM = -1.0

func main() {
  // runtime.GOMAXPROCS(1024)
  //Bring me the power of 1024 suns and an LG MEATS TEXAS STYLED BLT DRIPPING IN SOUTHERN STYLE STEAK SAUCE BROTHER
  dict.Initi("/mnt/c/Users/Maxim/go/src/dict/syllables")
  //shows where the syllables file is
  //TODO: variadic such that me and maxim don't have to swap it back and forth when either want to run it.

  dict.ToMap()
  //Maps are faster than array iteration, believe it or not.

  initi()
  //Initialization
  //prepares graph and samples

  go drawCostLoop()
  go drawGraphLoop(&nodeGraph)

  trainNetwork()
  cleanNetwork()
  manualTest()
}

func initi() {

  words = dict.SetOfKeys()
  sampleVariety := len(words)

  for i := 0; i < len(composition); i++ {
    nodeGraph[i] = make([]neuron, composition[i])
  }

  for i := 0; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].initNeuron(i,j)
    }
  }

  syllables = make([]int, sampleVariety)
  for i := 0; i < sampleVariety; i++ {
    //syllables[i] = dict.MapGet(words[i])
    syllables[i] = len(words[i])
  }

  forkCleanup(inputType)

}

func trainNetwork() {

  var firstCost float64 = 0.0
  var lastCost float64 = 0.0
  var generations int = 0
  var endTraining bool = false
  var minCostCheck int = 0

  for train := true; train; train = !endTraining {


    forkCycle(inputType)

    if generations == 0 {
      calcCost()
      fmt.Println("cost:",cost)
      firstCost = cost
    }

    backPropagation(totalSets)
    calcCost()
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
      manualTest();
    }

  }

  fmt.Println("gen", generations)

  calcCost()

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

func manualTest() {
  input := bufio.NewReader(os.Stdin)
  // inValue1 := 0.0
  // inValue2 := 0.0

  fmt.Println("Insert input")
  in,_ := input.ReadString('\n')
  for i:=0;i<len([]byte(in));i++{
    calcInputNeuron(i, joshRamp(float64([]byte(in)[i])), 0)
  }
  evaluateNetwork(0)

  for i := 0; i < composition[compLastRow]; i++ {
    fmt.Println("Output", (i + 1), ":", nodeGraph[compLastRow][i].RefInputSum[0])
  }

  manualTest()
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
