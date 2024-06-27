package main

import (
	"fmt"
	// "sort"
)

func main() {
	var a, b, c, t int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Printf("a = %d b = %d c = %d\n", a, b, c)
	if b > c {
		t = b
		b = c
		c = t	
	}
	if a > b {
		t = a
		a = b
		b = t
	}
	if b > c {
		t = b
		b = c
		c = t	
	}
	fmt.Printf("%d %d %d\n", a, b, c)
}


// func main() {
// 	var a, b, c int
// 	fmt.Scanf("%d %d %d", &a, &b, &c)
// 	list := []int{ a, b, c }
// 	sort.Slice( list, func(i, j int) bool { return list[i] < list[j]})
// 	fmt.Printf("%d %d %d\n", list[0], list[1], list[2])
// }