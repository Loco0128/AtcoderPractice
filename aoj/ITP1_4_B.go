package main

import "fmt"

func main (){
	var r float64
	pi := 3.141592653589
	fmt.Scanf("%f",&r)
	area := r * r * pi
	around := r * 2 * pi
	fmt.Printf("%f %f\n",area ,around)
}
