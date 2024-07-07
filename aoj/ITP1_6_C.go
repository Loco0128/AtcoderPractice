package main

import "fmt"

func main(){
	var n, b, f, r, v int
	var room [4][3][10]int

	fmt.Scan(&n)
	
	for i:=0;i<n;i++{
		fmt.Scanf("%d %d %d %d",&b, &f, &r, &v)
		room[b-1][f-1][r-1] += v 
	}

	for i,building := range room{
		for _,floor:=range building{
			for _,tenants:= range floor{
				fmt.Printf(" %d",tenants)
			}
			fmt.Printf("\n")
		}
		if i != 3{
			fmt.Printf("####################\n")
		}
	}
}