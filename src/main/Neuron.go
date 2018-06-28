package main

import
(
  "fmt"
  "math"
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
  neur.layer=layer
  neur.node=node
  // node:=neur.node

  neur.weights = make([]float64, composition[layer+1])
  neur.weightsChange = make([]float64, composition[layer+1])

  for i := 0; i < len(neur.weights); i++ {
    neur.weights[i] = rand.Float64()
  }
}

func (neur *neuron) calcInputSum() {
  fmt.Println("calcInputSum",neur.layer-1)
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
  // for i := 0; i<composition[0]; i++ {
  //   if i < 300 {
  //     nodeGraph[0][i].refInputSum = maxAmplitude[i]
  //   } else if i < 600 {
  //     nodeGraph[0][i].refInputSum = minAmplitude[i - 300]
  //   } else if i < 900 {
  //     nodeGraph[0][i].refInputSum = standardDeviationAmplitude[i - 600]
  //   } else if i < 1200 {
  //     nodeGraph[0][i].refInputSum = averageFrequency[i - 900]
  //   }
  // }
}

func sigmoid(input float64) float64{
  return 1 / (1 + (1/(math.Pow(math.E, input))))
}
