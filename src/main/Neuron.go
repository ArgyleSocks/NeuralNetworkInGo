package main

import
(
  //"fmt"
  "math"
  "time"
  "math/rand"
)

type neuron struct {
  layer int
  node int
  inputSum float64
  refInputSum float64
  outputSum float64
  weights []float64
  weightsChange []float64
}

func (neur *neuron) initNeuron(layer,node int) {
  neur.layer = layer
  neur.node = node
  // node:=neur.node

  if(layer != compLastRow) {
    neur.weights = make([]float64, composition[layer+1])
    neur.weightsChange = make([]float64, composition[layer+1])

    for i := 0; i < len(neur.weights); i++ {
      s1 := rand.NewSource(int64(time.Now().Nanosecond()))
      random := rand.New(s1)
      neur.weights[i] = random.Float64()
    }
  }
}

func (neur *neuron) calcInputSum() {
  //fmt.Println("calcInputSum",neur.layer-1)
  for i := 0; i < composition[neur.layer-1]; i++ {
    neur.inputSum += nodeGraph[neur.layer-1][i].calcOutputSum(neur.node)
  }
  neur.refInputSum = sigmoid(neur.inputSum)
}

func (neur *neuron) calcOutputSum(node int) float64{
  neur.outputSum=neur.refInputSum*neur.weights[node]
  return neur.outputSum
}

func calcInputNeuron() {//commented out because: A: dict first, B: the sound stuff isn't initialized and I needed to test.
  for i := 0; i<composition[0]; i++ {
    nodeGraph[0][i].refInputSum = 0.5
  }
}

func sigmoid(input float64) float64{
  return 1 / (1 + (1/(math.Pow(math.E, input))))
}
