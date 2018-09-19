package main

import
(
  "fmt"
  "math"
  "time"
  "math/rand"
)

type neuron struct {
  Layer int
  Node int
  InputSum []float64
  RefInputSum []float64
  OutputSum []float64
  Weights []float64
  TrainRel bool
  LocalDeriv float64
  WeightsChange []float64
}

func (neur *neuron) initNeuron(layer,node int) {

  neur.Layer = layer
  neur.Node = node
  neur.TrainRel = false
  neur.LocalDeriv = 0.0

  if(layer != compLastRow) {
    neur.Weights = make([]float64, composition[layer + 1])
    neur.WeightsChange = make([]float64, composition[layer + 1])

    for i := 0; i < len(neur.Weights); i++ {
      s1 := rand.NewSource(int64(time.Now().Nanosecond()))
      random := rand.New(s1)
      s2 := rand.NewSource(int64(time.Now().Nanosecond()))
      random2 := rand.New(s2)
      //fmt.Println(random.Float64(), random2.Float64())
      neur.Weights[i] = random.Float64() - random2.Float64()
    }
  }
}

func (neur *neuron) initSums(sets int) {
  fmt.Println("initedSums")
  neur.RefInputSum = make([]float64, sets)
  neur.InputSum = make([]float64, sets)
  neur.OutputSum = make([]float64, sets)

  for i := 0; i < composition[compLastRow]; i++ { //need to move this, like this really isn't supposed to be here
    expected[i] = make([]float64, totalSets)
  }
}

func (neur *neuron) calcInputSum(graph int) {
  //fmt.Println("calcInputSum",neur.layer-1)
  neur.InputSum[graph] = 0

  for i := 0; i < composition[neur.Layer-1]; i++ {
    neur.InputSum[graph] += nodeGraph[neur.Layer-1][i].calcOutputSum(neur.Node, graph)
  }
  neur.RefInputSum[graph] = sigmoid(neur.InputSum[graph])
}

func (neur *neuron) calcOutputSum(node int, graph int) float64{
  neur.OutputSum[graph] = neur.RefInputSum[graph] * neur.Weights[node]
  return neur.OutputSum[graph]
}

func calcInputNeuron(index int, input float64, set int) {
  //fmt.Println(len(nodeGraph[0][index].RefInputSum),set,index)
  nodeGraph[0][index].RefInputSum[set] = input
}

func sigmoid(input float64) float64{
  return 1 / (1 + (1/(math.Pow(math.E, input))))
}
