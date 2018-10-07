package main

import (
  "os"
  "fmt"
  // "time"
  "dict"
  "bufio"
  "runtime"
  // "strconv"
  // "strings"
  // "math/rand"
)

//TODO: Make sampleSet into 2 arrays with same length where one has string, other has syll

var composition[6]int = [...]int{16, 16, 16, 16, 16, 16}
var sampleSet [5][2]int = [...]int{{1, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2}}
var words []string //= dict.SetOfKeys
var syllables []int
var organizedWords [][]string
var organizedSyllables []int
var repValue = 5
var sampleVariety int //becomes len(words)
// var sampleSet [4][2]int = [4][2]int{{1, 1}, {2, 1}, {3, 1}, {4, 1}}
//var sampleSet map[string]int

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

var repetitionValue int = 50
var previousCost float64 = 0
var minimumCheck int
var endTraining bool = false

var numSamplesToTrain int = 10

var sampleVariableThingWeNeedToGetRidOfThis int = 0 //We need to seriously organize and also get rid of some of these global variables

func main() {
  runtime.GOMAXPROCS(1024)
  dict.Initi("/home/wurst/go/src/dict/syllables")
  dict.ToMap()
  initExpected(len(words)) //Need to move this to ExecNetwork, make it cycle and create additional nodeGraphs
  initi()
  go drawCostLoop()
  go drawGraphLoop(&nodeGraph)
  /*set:=dict.SetOfKeys()
  loset:=float64(len(set)) */
  //random (fast):
  /*for i:=0;i<numSamplesToTrain;i++ {
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

  fmt.Println("preparing training")
  prepareTraining()
  fmt.Println("done with preparation")
  tick()
  fmt.Println("tick")
  wait()
  fmt.Println("done")*/
  //cleanSamples()
  trainNetwork()
  cleanNetwork()
  manualTest()
}

func initi() {

  words = dict.SetOfKeys()
  sampleVariety = len(words)

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
    syllables[i] = dict.MapGet(words[i])
  }

  cleanSamples(2)

}

func trainNetwork() {

  //calcInputNeuron() Need to make this cycle through options in correspondence with the other thing

  /*TEST CODE:
  for i:=1;i<len(composition)-1;i++{
    for j:=0;j<composition[i];j++{
      //fmt.Println(nodeGraph[i][j].layer,nodeGraph[i][j].node)
    }
  }
  END*/

  for train := true; train; train = !endTraining {
    /*
    for i := 0; i < len(corresSet); i++ {
      for j := 0; j < corresSet[i][1]; j++ {
        //fmt.Println("j", j)
        setSample(corresSet[i][0], sampleVariableThingWeNeedToGetRidOfThis)
        evaluateNetwork(sampleVariableThingWeNeedToGetRidOfThis)
        sampleVariableThingWeNeedToGetRidOfThis++
      }
    }
    */
    for i := 0; i < len(organizedWords); i++ {
      for k := 0; k < repValue; k++ {
        //alright, so I'm doing this with the assumption sampleVariableThingWeNeedToGetRidOfThis corresponds to the set number, and i*k = the total number of elements per set, where i is every column in organized words, and k is in correspondence with the number of random words to pick.
        setSample(i,sampleVariableThingWeNeedToGetRidOfThis)
        evaluateNetwork(sampleVariableThingWeNeedToGetRidOfThis)
        sampleVariableThingWeNeedToGetRidOfThis++
      }
    }
    if generations == 0 {
      calcCost()
      fmt.Println("cost:",cost)
      firstCost = cost
    }

    backPropagation(totalSets) //make this sampleSet once you have added the cycle thing
    calcCost()
    if (cost == lastCost) || stableWeight {
      minimumCheck++
      if minimumCheck >= repetitionValue {
        endTraining = true
      }
    } else {
      minimumCheck = 0
    }

    // fmt.Println("COST:", cost, "CHANGE:", (cost - lastCost), "GENERATION:", generations)

    lastCost = cost
    generations++

    sampleVariableThingWeNeedToGetRidOfThis = 0
  }

  fmt.Println("gen", generations)

  /*for i := 0; i < composition[compLastRow]; i++ {
    //fmt.Println("Cost Node", (i + 1), "Expected:", expected[i], "Actual:", nodeGraph[compLastRow][i].refInputSum)
    if nodeGraph[compLastRow][i].RefInputSum > output {
      output = nodeGraph[compLastRow][i].RefInputSum
    }
  }

  //fmt.Print("Number of syllables: ")

  for i := 0; i < composition[compLastRow]; i++ {
    if nodeGraph[compLastRow][i].RefInputSum == output {
      //fmt.Println(i + 1)
    }
  }*/

  calcCost()

  fmt.Println("First cost:", firstCost, "\b, Last cost:", lastCost)
  fmt.Println("Change in cost:", (lastCost - firstCost) )
  //cleanup
  endTraining = false
  generations = 0
  minimumCheck = 0
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
    calcInputNeuron(i, float64([]byte(in)[i]), 0)
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
