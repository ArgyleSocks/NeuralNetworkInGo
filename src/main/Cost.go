package main

import (
  "fmt"
  "math"
  "dict"
  "time"
  // "reflect"
  "math/rand"
)

var compLastRow int = len(composition) - 1 //do we really need this A: replace all instances of compLastRow with len(composition) - 1 B: vice versa
var expected [][]float64 = make([][]float64, composition[compLastRow])
var cost float64 //doesn't need to global

func setSample(setIndex int, set int) {//finds a random value and puts it into the network
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
    calcInputNeuron(i, ramp(float64([]byte(k)[i]),45,122,-1,1), set)
  }
}

func initExpected(num int) {//supposed to set expected, but was converted to do the job setSample really does, but since setSample does it, it is obsolete. We still need to set expected.
  //sampleSet=make(map[string]int)

  //var word []byte

	for i:=0;i<num;i++{
		s1 := rand.NewSource(int64(time.Now().Nanosecond()))
    random := rand.New(s1)
    word:=dict.SetOfKeys()[int(random.Float64()*float64(len(dict.SetOfKeys())))]
    //index:=dict.MapGet(string(word))
    index:=len(word)
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
