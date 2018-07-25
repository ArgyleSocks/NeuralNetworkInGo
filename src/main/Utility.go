package main

import (
	// "io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

var first bool=true
var drawFile,_=os.OpenFile("drawBuffer.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)
var costFile,_=os.OpenFile("costBuffer.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)
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
		checkError(err)
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
	checkError(err)
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
