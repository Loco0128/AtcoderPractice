package main

import (
	"fmt"
	"math"
)

func main(){
	var a, b, c float64
	fmt.Scanf("%f %f %f",&a, &b, &c)
	s := a * b * calcSin(c) / 2
	l := math.Sqrt( math.Pow(a,2) + math.Pow(b,2) - 2 * a * b * calcCos(c)) + a + b
	h := s * 2 / a
	fmt.Printf("%.8f\n",s)
	fmt.Printf("%.8f\n",l)
	fmt.Printf("%.8f\n",h)

}

func calcSin(angle float64) float64{
	radians := angle * math.Pi / 180.0
	return math.Sin(radians)
}

func calcCos(angle float64) float64{
	radians := angle * math.Pi / 180.0
	return math.Cos(radians)
}
