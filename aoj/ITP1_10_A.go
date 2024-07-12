package main

import (
	"fmt"
	"math"
)

func main(){
	var p1 [2]float64
	var p2 [2]float64
	var ans float64
	fmt.Scanf("%f %f %f %f",&p1[0],&p1[1],&p2[0],&p2[1])
	pow := math.Pow((p2[0]-p1[0]),2)+math.Pow((p2[1]-p1[1]),2)
	ans = math.Sqrt(pow)
	fmt.Printf("%.8f\n",ans)
}
