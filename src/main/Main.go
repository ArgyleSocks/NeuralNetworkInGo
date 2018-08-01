package main

import (
  "fmt"
  "time"
  "dict"
  "runtime"
  "math/rand"
)

var composition[6]int=[...]int{20, 20, 20, 20, 20, 20}
var nodeGraph [][]neuron = make([][]neuron, len(composition))

var maxAmplitude []float64 //assuming 0 amplitude and the max overall amplitude corres to 0 and 1, adjust the values to be between 0 and 1 proportionally
var minAmplitude []float64
var standardDeviationAmplitude []float64 //see above
var averageFrequency []float64 //assuming 1 hz is 0.5 and the highest overall frequency is 1, adjust proportionally

var word []byte
var generations int = 0
var output float64 = 0

var firstCost float64
var lastCost float64

var repetitionValue int = 10
var previousCost float64 = 0
var minimumCheck int
var endTraining bool = false

var numSamplesToTrain int=150

func main() {
  runtime.GOMAXPROCS(1024)
  dict.Initi("/home/wurst/go/src/dict/syllables")
  dict.ToMap()
  initExpected()
  initi()
  go drawCostLoop()
  go drawGraphLoop(&nodeGraph)
  set:=dict.SetOfKeys()
  loset:=float64(len(set))
  //random (fast):
  for i:=0;i<numSamplesToTrain;i++ {
    ran:=rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
    word=[]byte(set[int(ran.Float64()*loset)])
    indexExpect:=dict.MapGet(string(word))
    for i2:=0;i2<len(expected);i2++ {
      if i2==indexExpect {
        expected[i2]=1
      } else {
        expected[i2]=0
      }
    }
    execNetwork()
    fmt.Println("Done With Sample",i)
  }
  //iterative (longer, more thorough):
  // for i:=0;i<len(set);i++ {
  //   fTime:=time.Now()
  //   word=[]byte(set[i])
  //   fmt.Println(string(word),"IS THE WORD")
  //   indexExpect:=dict.MapGet(string(word))
  //   for i2:=0;i2<len(expected);i2++ {
  //     if i2==indexExpect {
  //       expected[i2]=1
  //     } else {
  //       expected[i2]=0
  //     }
  //   }
  //   execNetwork()
  //   fmt.Println("Time taken for sample",i,"\b:",time.Now().Sub(fTime))
  // }
  word=[]byte("eelookoo")
  evaluateNetwork()
  for i2:=0;i2<len(expected);i2++ {
    if i2==3 {
      expected[i2]=1
    } else {
      expected[i2]=0
    }
  }
  calcCost(true)
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
  calcCost(false)
  firstCost = cost

  for train := true; train; train = !endTraining {
    evaluateNetwork()
    backPropagation()
    calcCost(false)
    if (cost == lastCost) || stableWeight {
      minimumCheck++
      if minimumCheck >= repetitionValue {
        endTraining = true
      }
    } else {
      minimumCheck = 0
    }
    lastCost = cost
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

  calcCost(false)
  lastCost = cost

  fmt.Println("First cost:", firstCost, "\b, Last cost:", lastCost)
  fmt.Println("Change in cost:", (lastCost - firstCost) )
  //cleanup
  endTraining=false
  generations=0
  minimumCheck=0

}

func evaluateNetwork() {
  for i := 1; i < len(composition); i++ {
    for j := 0; j < composition[i]; j++ {
      nodeGraph[i][j].calcInputSum()
    }
  }
}
