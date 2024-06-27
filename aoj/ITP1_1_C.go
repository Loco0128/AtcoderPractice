package main

import "fmt"

func main () {
	var row, col int
	fmt.Scanf("%d %d",&row, &col)

	area := row * col
	len := row * 2 + col * 2
	fmt.Printf("%d %d\n", area, len)
}