package main
import (
	// "fmt"
)
// type proc struct {
// 	firstWeightLayer int
// 	firstNode int
// 	lastWeightLayer int
// 	lastNode int
// }
// var procList []proc
var done bool=true
var ready [][][]chan bool
var derivTestRange float64=.0001
var derivBeforeEnd float64=.05
var maxitcount=100000
var houstonsaysgo bool=false
func prepareTraining() {//start processes and divy the weights up "evenly" (not really possible, but possible to a degree.)
	//get weight count
	// weightCount:=0
	// for i:=1;i<len(composition);i++ {
	// 	weightCount+=composition[i]*composition[i-1]
	// }
	// fmt.Println("> dividing",procCount,"processes among",weightCount,"weights")
	// var perProc int = weightCount/procCount
	// previousLayer:=0
	// previousNode:=-1
	// for i:=0;i<procCount;i++{
	// 	procList[i]=new(proc)
	// 	procList[i].firstWeightLayer=previousLayer+1
	// 	procList[i].firstNode=previousNode+1
	// 	if perProc>composition[procList[i].firstWeightLayer]*composition[procList[i].firstWeightLayer-1]
	// }
	ready=make([][][]chan bool,len(composition))
	for i:=0;i<len(composition);i++{
		ready[i]=make([][]chan bool,composition[i])
		for i2:=0;i2<composition[i];i2++{
			ready[i][i2]=make([]chan bool,len(nodeGraph[i][i2].Weights))
		}
	}
	for i:=0;i<len(composition);i++{
		for j:=0;j<composition[i];j++ {
			handleNode(i,j)
		}
	}
}
func tick() {
	houstonsaysgo=true
}
func wait() {
	for i:=0;i<len(composition);i++{
		for i2:=0;i2<composition[i];i2++ {
			for w:=0;w<len(nodeGraph[i][i2].Weights);w++{
				<-ready[i][i2][w]
			}
		}
	}
}
func handleNode(layer,node int) {
	for i:=0;i<len(nodeGraph[layer][node].Weights);i++{
		go handleWeight(layer,node,i)
	}
}
func handleWeight(l,n,w int) {
	for {
		if houstonsaysgo {
			//optimization...
			oldweight:=0.0
			oldcost:=0.0
			first:=true
			newweight:=0.0
			newcost:=0.0
			iterator:=0
			for {
				if iterator>=maxitcount{
					break
				}
				if first{
					oldweight=nodeGraph[l][n].Weights[w]
					evaluateNetwork()
					calcCost(true)//we now have x-y coordinates in the form of oldweight, oldcost
					oldcost=cost
					first=false
				} else {
					oldweight=nodeGraph[l][n].Weights[w]
					oldcost=newcost
				}
				nodeGraph[l][n].Weights[w]-=derivTestRange
				newweight=nodeGraph[l][n].Weights[w]
				evaluateNetwork()
				calcCost(true)
				newcost=cost
				deriv:=(newcost-oldcost)/(newweight-oldweight)
				nodeGraph[l][n].Weights[w]-=deriv
				iterator++
			}
			ready[l][n][w]<-true
		}
	}
}