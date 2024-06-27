package main

import "fmt"

func main() {
	var w,h,x,y,r int 
	
	fmt.Scanf("%d %d %d %d %d", &w, &h, &x, &y, &r)

	if x + r > w || x - r < 0 || y + r > h || y - r < 0 {
		fmt.Printf("No\n")
	} else {
		fmt.Printf("Yes\n")
	}
}
