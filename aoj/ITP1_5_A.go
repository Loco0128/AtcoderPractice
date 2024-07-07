package main

import "fmt"

func main(){
	var h, w int
	var input [][2]int
	for {
		fmt.Scanf("%d %d", &h, &w)
		if h == 0 && w == 0{
			break
		}
		input  = append(input, [2]int{h,w})
	}
	for _,v := range input {
		printFigure(v[0],v[1])
	}
}

func printFigure(h,w int){
	for i:=0;i<h;i++{
		for j:=0;j<w;j++{
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}