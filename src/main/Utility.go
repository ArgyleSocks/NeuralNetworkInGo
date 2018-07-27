package main

import (
	// "io/ioutil"
	"encoding/json"
	"math"
	"fmt"
	"os"
)

var first bool=true
var drawFile,_=os.OpenFile("drawBuffer.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)
var costFile,_=os.OpenFile("costBuffer.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)

func checkNaN(val float64) {
	if math.IsNaN(val) {
		fmt.Fprintf(os.Stderr, "An NaN: %s\n", val)
        os.Exit(1)
	}
}

func drawGraphLoop(graph *[][]neuron){
	for {
		drawGraph(*graph)
	}
}
func drawCostLoop(){
	arr:=make([]float64,50)
	for {
		for e:=0;e<len(arr);e++{
			if e<len(arr)-1{
				arr[e]=arr[e+1]
			} else {
				arr[e]=cost
			}
		}
		j,err:=json.Marshal(arr)
		// fmt.Println("Little tim needs his revolver", arr)
		// checkError(err)
		j=[]byte(string(j)+"end")
		err=costFile.Truncate(0)
		checkError(err)
		_,err=costFile.WriteAt([]byte(j),0)
		checkError(err)
	}
}
func drawGraph(graph [][]neuron) {//draw nodeGraph
	if first{
		fmt.Println("make sure you start renderGraph.py if you want graphics (nuklear is awful)")
		first=false
	}
	j,err:=json.Marshal(graph)
	// fmt.Println("Little tim needs his linux handbook")
	// checkError(err)
	j=[]byte(string(j)+"end")
	// fmt.Println(graph)
	err=drawFile.Truncate(0)
	checkError(err)
	_,err=drawFile.WriteAt([]byte(j),0)
	checkError(err)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "An error: %s\n", err.Error())
        os.Exit(1)
    }
}
