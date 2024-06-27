package main

import "fmt"

func main () {
	var time int
	fmt.Scanf("%d",&time)

	hour := time / 60 / 60
	min := time / 60 - hour * 60
	sec := time  - min * 60 - hour * 60 * 60
	fmt.Printf("%d:%d:%d\n", hour, min, sec)
}