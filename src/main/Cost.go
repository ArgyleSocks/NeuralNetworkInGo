package main

import (
  "fmt"
  // "math"
  "dict"
  "time"
  // "reflect"
  "math/rand"
)

var compLastRow int = len(composition) - 1 //Turns out this helps in context
var expected [][]float64 = make([][]float64, composition[compLastRow])
var cost float64 //doesn't need to global

func initExpected(expectedResult []float64, set int) {//supposed to set expected, but was converted to do the job setSample really does, but since setSample does it, it is obsolete. We still need to set expected.
  for i, e := range expectedResult {
    expected[i][set] = 0
    expected[i][set] = e
  }
  //expected=append(expected,expectedSampleResult)  OBSOLETE
}

func setSyllableCountAndWordRandomFromDictLikeABossBoiii(num int) {
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
