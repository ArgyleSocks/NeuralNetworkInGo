package main

import (
  "math/rand"
  "strconv"
  "strings"
  "bufio"
  "math"
  "time"
  "fmt"
  "os"
)
//Utility? -->
//Possibly Irrelevant/Needs to be implemented

func calcCost(manual bool) {
  cost = 0
  for i := 0; i < len(expected); i++ {
    if !manual {
      for j := 0; j < len(expected[0]); j++ {
        cost += math.Pow((nodeGraph[compLastRow][i].RefInputSum[j] - expected[i][j]), 2)
        checkNaN(cost)
      }
    } else {
      cost += math.Pow((nodeGraph[compLastRow][i].RefInputSum[0] - expected[i][0]), 2)
      checkNaN(cost)
    }
  }
}

func nodeInputSum(layer int, node int, set int) float64 {
  return nodeGraph[layer][node].InputSum[set]
}

//Possibly Irrelevant/Needs to be implemented
func nodeRefInputSum(layer int, node int, set int) float64 {
  return nodeGraph[layer][node].RefInputSum[set]
}

//Possibly Irrelevant/Needs to be implemented
func nodeWeight(layer int, node int, corresNode int) float64 {
  return nodeGraph[layer][node].Weights[corresNode]
}

func forkRefInputSum(refInputSumType int, input float64) float64 {
  switch refInputSumType {
  case 1:
    return sigmoid(input)
  case 2:
    return ramp(input)
  case 3:
    return joshRamp(input)
  case 4:
    //Nothing yet!
  }
  fmt.Println("Not a valid input for forkRefInputSum")
  return 0
}

func forkDerivative(refInputSumType int, input float64) float64 {
  switch refInputSumType {
  case 1:
    return sigmoidDerivative(input)
  case 2:
    return rampDerivative(input)
  case 3:
    return joshRampDerivative(input)
  case 4:
    //Nothing yet!
  }
  fmt.Println("Not a valid input for forkDerivative")
  return 0
}

func sigmoid(input float64) float64 {
  return 1 / (1 + (1/(math.Pow(math.E, input))))
}

func sigmoidDerivative(input float64) float64 {
  return (1/(math.Pow((1 + math.Pow(math.E, -input)), 2) * math.Pow(math.E, input)))
}

func trainingRate(input float64) float64 {
  return (2 * sigmoidDerivative(0.25 * input) + 0.02)
}

func ramp(input float64) float64 {
  if input > 0 {
    return input
  } else {
    return 0
  }
}

func rampDerivative(input float64) float64 {
  if input > 0 {
    return 1
  } else {
    return 0
  }
}

func joshRamp(input float64) float64 {
	currentRange := UPPER_LIM - LOWER_LIM
	desiredRange := DESIRED_UPPER_LIM - DESIRED_LOWER_LIM
	point := input/currentRange
	return point * desiredRange + DESIRED_LOWER_LIM
}

func joshRampDerivative(input float64) float64 {
  currentRange := UPPER_LIM - LOWER_LIM
	desiredRange := DESIRED_UPPER_LIM - DESIRED_LOWER_LIM
  return desiredRange / currentRange
}

func forkCleanup(sampleType int) {
  switch sampleType {
  case 1:
    twoDiCleanup()
  case 2:
    uniformCasesCleanup()
  case 3:
    //Nothing yet!
  }
}

func forkCycle(sampleType int) {
  switch sampleType {
  case 1:
    fmt.Println("You are using two di")
    twoDiCycle()
  case 2:
    fmt.Println("You are using uniform")
    uniformCasesCycle()
  case 3:
    //Nothing yet!
  }
}

func forkTrainingTask(trainingTask int, setType int, setIndex int) {
  switch trainingTask {
  case 1:
    if sampleType == 1 {
      logicPuzzleTraining(setType, setIndex)
    } else {
      fmt.Println("Invalid sampleType for this trainingTask")
    }
  case 2:
    if sampleType == 2 {
      letterCountTraining(setType, setIndex)
    } else {
      fmt.Println("Invalid sampleType for this trainingTask")
    }
  case 3:
    if sampleType == 2 {
      //WIP
    } else {
      fmt.Println("Invalid sampleType for this trainingTask")
    }
  case 4:
    //Nothing yet!
  }
}

func logicPuzzleTraining(setType int, setIndex int) {
  switch setType {
  case 1:
    var wanted []float64 = []float64{0,0}//instead of setting these here, reference variables from a TODO training object which stores data/expected, trainingRate, etc. Allows for modularity.
    var input []float64 = []float64{0,0}
    calcInputNeuron(input, setIndex)//change to match initExpected's arguments
    initExpected(wanted, setIndex)
  case 2:
    var wanted []float64 = []float64{1,0}
    var input []float64 = []float64{1,0}
    calcInputNeuron(input, setIndex)
    initExpected(wanted, setIndex)
  case 3:
    var wanted []float64 = []float64{0,1}
    var input []float64 = []float64{0,1}
    calcInputNeuron(input, setIndex)
    initExpected(wanted, setIndex)
  case 4:
    var wanted []float64 = []float64{0,0}
    var input []float64 = []float64{1,1}
    calcInputNeuron(input, setIndex)
    initExpected(wanted, setIndex)
  }
}

func letterCountTraining(setType int, setIndex int) {
  s1 := rand.NewSource(int64(time.Now().Nanosecond()))
  random := rand.New(s1)
  //index finding
  k := organizedWords[setType][int(random.Float64()*(float64(len(organizedWords[setIndex])))]
  kByteArray := []byte(k)
  kFloatArray := make([]float64, composition[0])

  for i := 0; i < composition[0]; i++ {
    if i < len(kByteArray) {
      kFloatArray[i] = joshRamp(kByteArray[i])
    } else {
      kFloatArray[i] = 0
    }
  }
  //iterate through, set input layer accordingly
  for i := 0 ; i < len(k) ; i++ {
    calcInputNeuron(, setIndex)
  }// TODO: implement calcInputNeuron/initExpected*/
}

func arithmeticTraining(set int, setIndex int) {
  //Nothing Yet!
}

func forkManualTest(trainingTask int) {
  switch trainingTask {
  case 1:
    manualTestFloatInput();
  case 2:
    manualTestAsciiInput();
  case 3:
    //WIP
  case 4:
    //Nothing Yet!
  }

}

func manualTestAsciiInput() {
  input := bufio.NewReader(os.Stdin)

  fmt.Println("Insert input string:")
  in,_ := input.ReadString('\n')
  inputArr:=make([]float64,0)
  for i := 0; i < len([]byte(in)); i++{
    inputArr=append(inputArr,joshRamp(float64([]byte(in)[i])))
  }
  calcInputNeuron(inputArr, 0)
  //uses the nodeGraph at graph 0
  evaluateNetwork(0)

  for i := 0; i < composition[compLastRow]; i++ {
    fmt.Println("Output", (i + 1), ":", nodeGraph[compLastRow][i].RefInputSum[0])
  }

  manualTestAsciiInput()
}

func manualTestFloatInput() {
  input := bufio.NewReader(os.Stdin)
  expectedInput := bufio.NewReader(os.Stdin)
  fmt.Println("Input float/int values on separated solely by commas:")
  inputArr := make([]float64, 0)
  expectedArr := make([]float64, 0)
  for i, _ := range nodeGraph[0] {
    fmt.Println((i + 1),"th input")
    string, err := input.ReadString('\n')
    fmt.Println(err)
    value, err := strconv.ParseFloat(strings.Split(string,"\n")[0], 64)
    fmt.Println(err)
    inputArr = append(inputArr,value)
  }

  fmt.Println("Input the expected")
  for i, _ := range expected {
    fmt.Println((i + 1),"th expected")
    string, err := expectedInput.ReadString('\n')
    fmt.Println(err)
    value, err := strconv.ParseFloat(strings.Split(string,"\n")[0], 64)
    fmt.Println(err)
    expectedArr = append(expectedArr, value)
  }

  initExpected(expectedArr, 0)
  fmt.Println(inputArr)
  calcInputNeuron(inputArr, 0)

  evaluateNetwork(0)
  calcCost(true)
  fmt.Println("Cost:", cost)

  for i := 0; i < composition[compLastRow]; i++ {
    fmt.Println("Output", (i + 1), ":", nodeGraph[compLastRow][i].RefInputSum[0], ", Expected:", expected[i][0])
  }

  manualTestFloatInput()
}
