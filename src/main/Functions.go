package main

import (
  "strconv"
  "strings"
  "bufio"
  "math"
  "fmt"
  "os"
)
//Utility? -->
//Possibly Irrelevant/Needs to be implemented

func (net *network) calcCost(manual bool) {
  cost = 0
  for i := 0; i < len(expected); i++ {
    if !manual {
      for j := 0; j < len(expected[0]); j++ {
        //fmt.Println("ex",expected)
        checkNaN(net.NodeGraph[compLastRow][i].RefInputSum[j])
        cost += math.Pow((net.NodeGraph[compLastRow][i].RefInputSum[j] - expected[i][j]), 2)
        checkNaN(cost)
      }
    } else {
      cost += math.Pow((net.NodeGraph[compLastRow][i].RefInputSum[0] - expected[i][0]), 2)
      checkNaN(cost)
    }
  }
}

func (net *network) nodeInputSum(layer int, node int, set int) float64 {
  return net.NodeGraph[layer][node].InputSum[set]
}

//Possibly Irrelevant/Needs to be implemented
func (net *network) nodeRefInputSum(layer int, node int, set int) float64 {
  return net.NodeGraph[layer][node].RefInputSum[set]
}

//Possibly Irrelevant/Needs to be implemented
func (net *network) nodeWeight(layer int, node int, corresNode int) float64 {
  return net.NodeGraph[layer][node].Weights[corresNode]
}

func forkRefInputSum(refInputSumType int, input float64) float64 {
  switch refInputSumType {
  case 1:
    return sigmoid(input)
  case 2:
    return ramp(input)
  case 3:
    //return joshRamp(input)
  case 4:
    //Nothing yet!
  }
  fmt.Println("Not a valid input for forkRefInputSum")
  return 0
}

func forkDerivative(refInputSumType int, input float64) float64 {
  switch refInputSumType {
  case 1:
    checkNaN(sigmoidDerivative(input))
    return sigmoidDerivative(input)
  case 2:
    return rampDerivative(input)
  case 3:
    //return joshRampDerivative(input)
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
  checkNaN(input)
  if math.Pow(math.E, input) <= 0 {
    return 0
  } else {
    return (1/(math.Pow((1 + math.Pow(math.E, -input)), 2) * math.Pow(math.E, input)))
  }
}

func trainingRate(input float64) float64 {
  return (2 * sigmoidDerivative(.5 * input) + 0.02)
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

func asciiCompression(input float64) float64 {
	yRange := DESIRED_UPPER_ASCII_LIM - DESIRED_LOWER_ASCII_LIM
	xRange := UPPER_ASCII_LIM - LOWER_ASCII_LIM
	slope := yRange/xRange
	return (input - LOWER_ASCII_LIM) * slope + DESIRED_LOWER_ASCII_LIM
}

/*func joshRampDerivative(input float64) float64 {
  currentRange := UPPER_ASCII_LIM - LOWER_ASCII_LIM
	desiredRange := DESIRED_UPPER_INPUT_LIM - DESIRED_LOWER_INPUT_LIM
  return desiredRange / currentRange
}*/

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
    //fmt.Println("You are using two di")
    twoDiCycle()
  case 2:
    //fmt.Println("You are using uniform")
    uniformCasesCycle()
  case 3:
    //Nothing yet!
  }
}

func (net *network) logicPuzzleTraining(setType int, setIndex int) {
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

func arithmeticTraining(set int, setIndex int) {
  //Nothing Yet!
}

func (net *network) forkManualTest(trainingTask int) {
  switch trainingTask {
  case 1:
    net.manualTestFloatInput();
  case 2:
    net.manualTestAsciiInput();
  case 3:
    //WIP
  case 4:
    //Nothing Yet!
  }

}

func (net *network) manualTestAsciiInput() {
  input := bufio.NewReader(os.Stdin)

  fmt.Println("Insert input string:")
  in,_ := input.ReadString('\n')
  inputArr:=make([]float64,len(net.NodeGraph[0]))
  for i := 0; i < len([]byte(in)); i++{
    inputArr[i]=asciiCompression(float64([]byte(in)[i]))
  }
  for i := 0; i < len(inputArr); i++ {
    if inputArr[i]==0 {
      inputArr[i]=-999
    }
  }
  fmt.Println(inputArr)
  calcInputNeuron(inputArr, 0)
  //uses the net.NodeGraph at graph 0
  evaluateNetwork(0)

  for i := 0; i < net.Composition[compLastRow]; i++ {
    fmt.Println("Output", (i + 1), ":", net.NodeGraph[compLastRow][i].RefInputSum[0])
  }

  net.manualTestAsciiInput()
}

func (net *network) manualTestFloatInput() {
  input := bufio.NewReader(os.Stdin)
  expectedInput := bufio.NewReader(os.Stdin)
  fmt.Println("Input float/int values on separated solely by commas:")
  inputArr := make([]float64, 0)
  expectedArr := make([]float64, 0)
  for i, _ := range net.NodeGraph[0] {
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
  net.calcCost(true)
  fmt.Println("Cost:", cost)

  // for i := 0; i < composition[compLastRow]; i++ {
  //   fmt.Println("Output", (i + 1), ":", net.NodeGraph[compLastRow][i].RefInputSum[0], ", Expected:", expected[i][0])
  // }

  net.manualTestFloatInput()
}

func (net *network) manualTest(inputArray []float64, expectedArray []float64) {
  fmt.Println("normal ManTest")
  calcInputNeuron(inputArray, 0)
  initExpected(expectedArray, 0)

  evaluateNetwork(0)
  net.calcCost(true)
  fmt.Println("Cost:", cost)
}
