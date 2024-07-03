package main

import "fmt"

func main (){
	var a, b int

	fmt.Scanf("%d %d %d %d",&a ,&b)
	d := a / b
	r := a % b
	f := float64(a) / float64(b)
	fmt.Printf("%d %d %f\n",d ,r ,f)
}
