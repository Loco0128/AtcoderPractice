 package main

 import "fmt"

 func main(){
	var n,x int
	fmt.Scanf("%d",&n)
	for i:=1;i <=n;i++{
		x = i
		if x % 3 == 0{
			fmt.Printf(" %d", i)
		} else {
			for x != 0 {
				if x % 10 == 3{
					fmt.Printf(" %d", i)
					break
				} else {
					x = x / 10
				}
			}
		}
	}
	fmt.Printf("\n")
 }