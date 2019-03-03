package main

import
(
  //"fmt"
  //"math"
  "time"
  "math/rand"
)

type neuron struct {
  Layer int
  Node int
  InputSum []float64//don't worry about this unless changing sigmoid/ReLu/whatever the individual node function is.
  RefInputSum []float64//what to output to other nodes
  OutputSum []float64
  Weights []float64
  TrainRel bool
  LocalDeriv float64
  WeightsChange []float64
}

func (neur *neuron) initNeuron(layer, node int) {

  neur.Layer = layer
  neur.Node = node
  //knows where it is in the graph

  neur.TrainRel = false
  //am I relevant?

  neur.LocalDeriv = 0.0
  //Amount to adjust the weight along the dimension that this node represents

  if(layer != compLastRow) {
    neur.Weights = make([]float64, composition[layer + 1])
    //what do you think?

    neur.WeightsChange = make([]float64, composition[layer + 1])
    //amount that weights are scheduled to be changed by

    for i := 0; i < len(neur.Weights); i++ {
      s1 := rand.NewSource(int64(time.Now().Nanosecond()))
      random := rand.New(s1)
      s2 := rand.NewSource(int64(time.Now().Nanosecond()))
      random2 := rand.New(s2)
      //fmt.Println(random.Float64(), random2.Float64())
      neur.Weights[i] = random.Float64() - random2.Float64()
    }
    //we just set all of the weights to a random value, nuff' said
  }
}

func (neur *neuron) initSums(sets int) { //has been replaced in InitNetworkVar
  //fmt.Println("initSums",sets)
  neur.RefInputSum = make([]float64, sets)
  neur.InputSum = make([]float64, sets)
  neur.OutputSum = make([]float64, sets)
}

func (neur *neuron) calcInputSum(graph int) {
  //find what to output based on inputs.

  neur.InputSum[graph] = 0

  for i := 0; i < composition[neur.Layer - 1]; i++ {
    neur.InputSum[graph] += nodeGraph[neur.Layer - 1][i].calcOutputSum(neur.Node, graph)
  }
  neur.RefInputSum[graph] = forkRefInputSum(refInputSumType, neur.InputSum[graph])
}

func (neur *neuron) calcOutputSum(node int, graph int) float64{
  neur.OutputSum[graph] = neur.RefInputSum[graph] * neur.Weights[node]
  return neur.OutputSum[graph]
}

func calcInputNeuron(input []float64, setIndex int) {
  //fmt.Println("input array:", input, "setIndex:", setIndex)
  for i, e := range input {
    nodeGraph[0][i].RefInputSum[setIndex] = e
  }
}
