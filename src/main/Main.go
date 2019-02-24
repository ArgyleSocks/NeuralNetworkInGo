package main

import (
  "fmt"
  "github.com/thesupremeliverwurst/dict"
  "time"
  "math/rand"
  "math"
)

/*var LPComposition []int = []int{2, 3, 3, 3, 3, 2}
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
}*/

var genericRepeat int = 1

var LPComposition []int = []int{12, 20, 30, 20, 13, 13}

var numTrainFor int = 75
var maxWordLength int = 12

var trainingSet1 [][2]int//auto-gen

var inputDataSet1 [][]float64//auto-gen
var expectedDataSet1 [][]float64

func main() {
  fmt.Println("I Feel Special")
  //GENERATION
  dict.Initi("../../dat/syllables")
  dict.ToMap()
  keys := dict.SetOfKeys()
  //set up trainingSet1
  trainingSet1 = make([][2]int, int(math.Abs(float64(numTrainFor))))
  if numTrainFor < 0 && maxWordLength == 0 {
    fmt.Println("Using no. 1")
    for i := 0; i > numTrainFor; i-- {
      trainingSet1[-i][0] = -i
      trainingSet1[-i][1] = genericRepeat//lazy approach
    }
  } else if numTrainFor < 0 {
    trainCounter := 0
    for i := 0; i < -numTrainFor; i++ {
      fmt.Println("Using no. 2")
      if dict.MapGet(keys[i]) <= maxWordLength {
        fmt.Println(trainCounter, len(keys))
        trainingSet1[trainCounter][0] = i
        trainingSet1[trainCounter][1] = genericRepeat//lazy approach
        trainCounter++
      }

      if trainCounter > -numTrainFor {
        break
      }
    }
  } else {
    for i := 0; i < numTrainFor; i++ {
      fmt.Println("Using no. 3")
      s := rand.NewSource(int64(time.Now().Nanosecond()))
      random := rand.New(s)
      index := int(random.Float64()*float64(len(keys)))
      trainingSet1[i][0] = index
      trainingSet1[i][1] = genericRepeat//lazy approach
    }
  }
  //set up input/output
  inputDataSet1=make([][]float64,len(keys))
  expectedDataSet1=make([][]float64,len(keys))
  for i := 0; i < len(keys); i++ {
    inputDataSet1[i]=make([]float64,LPComposition[0])
    bArr:=[]byte(keys[i])
    for i2 := 0; i2 < LPComposition[0]; i2++ {
      if i2 < len(bArr) {
        inputDataSet1[i][i2]=joshRamp(float64(bArr[i2]))//TODO move joshRamp variables here or fix joshRamp to take arguments again... ahem.
      } else {
        inputDataSet1[i][i2]=-2
      }
    }
    expectedDataSet1[i]=make([]float64,LPComposition[len(LPComposition)-1])
    //find length of word
    rIndex:=0
    for ch:=0; ch<len(inputDataSet1[i]); ch++ {
        //fmt.Println(inputDataSet1[i][ch])
        if inputDataSet1[i][ch]!=-2 {
            rIndex++
        } else {
            break
        }
    }
    expectedDataSet1[i][rIndex]=1
  }
  //END GENERATION
  //fmt.Println(inputDataSet1)
  fmt.Println()
  //fmt.Println(expectedDataSet1)
  //NETWORK
  InitNetworkVar(LPComposition, inputDataSet1, expectedDataSet1, trainingSet1)
  NeuralNetworkExec()
  //END NETWORK
}
//possibly useful in future?
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
