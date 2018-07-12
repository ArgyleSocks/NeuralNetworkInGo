package main

import
(
  //"fmt"
  "math"
  "time"
  "math/rand"
)

type neuron struct {
  Layer int
  Node int
  InputSum float64
  RefInputSum float64
  OutputSum float64
  Weights []float64
  WeightsChange []float64
}

func (neur *neuron) initNeuron(layer,node int) {
  neur.Layer=layer
  neur.Node=node
  // node:=neur.node

  if(layer != compLastRow) { 
    neur.Weights = make([]float64, composition[layer+1])
    neur.WeightsChange = make([]float64, composition[layer+1])

    for i := 0; i < len(neur.Weights); i++ {
      s1 := rand.NewSource(int64(time.Now().Nanosecond()))
      random := rand.New(s1)
      neur.Weights[i] = random.Float64()
    }
  }
}

func (neur *neuron) calcInputSum() {
  //fmt.Println("calcInputSum",neur.layer-1)
  for i := 0; i < composition[neur.Layer-1]; i++ {
    neur.InputSum += nodeGraph[neur.Layer-1][i].calcOutputSum(neur.Node)
  }
  neur.RefInputSum = sigmoid(neur.InputSum)
}

func (neur *neuron) calcOutputSum(node int) float64{
  neur.OutputSum=neur.RefInputSum*neur.Weights[node]
  return neur.OutputSum
}

func calcInputNeuron() {//commented out because: A: dict first, B: the sound stuff isn't initialized and I needed to test.
  for i := 0; i<composition[0]; i++ {
    nodeGraph[0][i].RefInputSum = 0.5
  }
}

func sigmoid(input float64) float64{
  return 1 / (1 + (1/(math.Pow(math.E, input))))
}
