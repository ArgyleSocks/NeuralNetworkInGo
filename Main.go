package main

import (
  "fmt"
)

var LPComposition []int = []int{2, 3, 3, 3, 3, 2}
var trainingSet1 [][2]int = [][2]int{{0, 1}, {1, 1}, {2, 1}, {3, 1}}
//The 1st index of trainingSet will correspond to a set index in inputDataSet.
//The 2nd index will correspond to the number of times it should be repeated.
//he even wrote it right here... how could I forget? 🎸🎸🎸🎸🎸🎸🎸♭
var inputDataSet1 [][]float64 = [][]float64{{0, 0}, {1, 0}, {0, 1}, {1, 1}}
var expectedDataSet1 [][]float64 = [][]float64{{0, 0}, {0, 1}, {1, 0}, {0, 0}}

func main() {
  fmt.Println("Program Started")
  InitNetworkVar(LPComposition, inputDataSet1, expectedDataSet1, trainingSet1)
  NeuralNetworkExec()
}
//possibly useful in future? yeah no don't worry about it; anything
//associated with dict/letterCount will be useful. Just rn it should be commented to narrow our focus
//to organization into a sampelArrangement-less/library friendly format
/*func setSyllableCountAndWordRandomFromDictLikeABossBoiii(num int) {//I feel crucified

  for i:=0;i<num;i++{
    s1 := rand.NewSource(int64(time.Now().Nanosecond()))
    random := rand.New(s1)
    word:=dict.SetOfKeys()[int(random.Float64()*float64(len(dict.SetOfKeys())))]
    //index:=dict.MapGet(string(word))
    index:=len(word)
    words[i]=word
    syllables[i]=index
    //trainingSet[word]=index
  }
  fmt.Println("Expected initialized")
}
*/
