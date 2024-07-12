package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	var score float64
	var allScores [][]float64
	for {
		var scores []float64
		fmt.Scan(&num)
		if num == 0{
			break
		}
		for i:=0;i<num;i++{
			fmt.Scan(&score)
			scores = append(scores, score)
		}
		allScores = append(allScores, scores)
	}

	for _,v := range allScores{
		fmt.Printf("%.8f\n",calc(v,len(v)))
	}
}

func calc(dataset []float64, leng int)float64{
	var sum float64
	ave := calcAve(dataset,leng)
	for _,v:= range dataset{
		sum += math.Pow(v-ave,2)
	}
	return math.Sqrt(sum / float64(leng))


}

func calcAve(dataset []float64, leng int)float64{
	var sum float64
	for _,v := range dataset{
		sum += v
	}
	return sum / float64(leng)
}
