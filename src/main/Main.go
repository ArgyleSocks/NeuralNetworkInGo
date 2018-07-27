package main

import (
  "fmt"
  "dict"
  "runtime"
)

var composition[6]int=[...]int{6, 6, 6, 6, 6, 6}
var nodeGraph [][]neuron = make([][]neuron, len(composition))

var maxAmplitude []float64 //assuming 0 amplitude and the max overall amplitude corres to 0 and 1, adjust the values to be between 0 and 1 proportionally
var minAmplitude []float64
var standardDeviationAmplitude []float64 //see above
var averageFrequency []float64 //assuming 1 hz is 0.5 and the highest overall frequency is 1, adjust proportionally

var word []byte
var generations int = 0
var output float64 = 0

func main() {
  runtime.GOMAXPROCS(1024)
  dict.Initi("/home/wurst/go/src/dict/syllables")
  dict.ToMap()
  word=[]byte(dict.SetOfKeys()[0])
  initExpected()
  initi()
  go drawCostLoop()
  go drawGraphLoop(&nodeGraph)
  execNetwork()
  /*
  fmt.Println("preparing training")
  prepareTraining()
  fmt.Println("done with preparation")
  tick()
  fmt.Println("tick")
  wait()
  fmt.Println("done")*/
}

func initi() {
  for i := 0; i < len(composition); i++ {
    nodeGraph[i] = make([]neuron, composition[i])
  }

  for i := 0; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].initNeuron(i,j)
    }
  }
}

func execNetwork() {

  calcInputNeuron()//prepare peripherals

  /*TEST CODE:
  for i:=1;i<len(composition)-1;i++{
    for j:=0;j<composition[i];j++{
      //fmt.Println(nodeGraph[i][j].layer,nodeGraph[i][j].node)
    }
  }
  END*/

  evaluateNetwork()
  calcCost()
  for train := true; train; train = !(endTraining && (generations > 100)) {
    evaluateNetwork()
    backPropagation()
    calcCost()
    generations++
  }

  fmt.Println("gen", generations)

  /*for i := 0; i < composition[compLastRow]; i++ {
    //fmt.Println("Cost Node", (i + 1), "Expected:", expected[i], "Actual:", nodeGraph[compLastRow][i].refInputSum)
    if nodeGraph[compLastRow][i].RefInputSum > output {
      output = nodeGraph[compLastRow][i].RefInputSum
      fmt.Println("checking output")
      checkNaN(output)
    }
  }

  //fmt.Print("Number of syllables: ")

  for i := 0; i < composition[compLastRow]; i++ {
    if nodeGraph[compLastRow][i].RefInputSum == output {
      //fmt.Println(i + 1)
    }
  }*/

  calcCost()

}

func evaluateNetwork() {
  for i := 1; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].calcInputSum()
    }
  }
}
