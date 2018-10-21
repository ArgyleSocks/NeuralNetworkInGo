package main

import (
  "fmt"
  "math"
  "dict"
  "time"
  // "reflect"
  "math/rand"
)

var compLastRow int = len(composition) - 1
var expected [][]float64 = make([][]float64, composition[compLastRow])
var cost float64
var costDeriv float64

 func setSample(setIndex int, set int) {
  /*v := reflect.ValueOf(reflect.ValueOf(sampleSet).MapKeys()).Interface().([]string)
  fmt.Println(len(v))
  fmt.Println("setValue", setValue)
  fmt.Println("refSum", len(nodeGraph[0][0].RefInputSum), "set", set)
  switch setValue {
  case 1:
    calcInputNeuron(0, 0, set)
    initExpected(0, 0, set)
  case 2:
    calcInputNeuron(1, 0, set)
    initExpected(1, 0, set)
  case 3:
    calcInputNeuron(0, 1, set)
    initExpected(0, 1, set)
  case 4:
    calcInputNeuron(1, 1, set)
    initExpected(0, 0, set)
  }*/

  // Let's find a word!

  //random seed
  s1 := rand.NewSource(int64(time.Now().Nanosecond()))
  random := rand.New(s1)

  //index finding
  k := organizedWords[setIndex][int(random.Float64()*float64(len(organizedWords[setIndex])))]

  //iterate through, set input layer accordingly
  //fmt.Println(k)
  for i := 0 ; i < len(k) ; i++ {
    calcInputNeuron(i, float64([]byte(k)[i]), set)
  }
}

func initExpected(num int) {
  //sampleSet=make(map[string]int)
	for i:=0;i<num;i++{
		s1 := rand.NewSource(int64(time.Now().Nanosecond()))
    random := rand.New(s1)
    word:=dict.SetOfKeys()[int(random.Float64()*float64(len(dict.SetOfKeys())))]
    index:=dict.MapGet(string(word))
    words[i]=word
    syllables[i]=index
    //sampleSet[word]=index
	}
  fmt.Println("Expected initialized")
}

func calcCost() {
  cost = 0
  for i := 0; i < len(expected); i++ {
    for j := 0; j < len(expected[0]); j++ {
      cost += math.Pow((nodeGraph[compLastRow][i].RefInputSum[j] - expected[i][j]), 2)
      checkNaN(cost)
    }
  }
}
