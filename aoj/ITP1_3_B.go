package main

import "fmt"

func main(){
	var input []int
	var num int

	for {
		fmt.Scanf("%d", &num)
		if num == 0 {
			break
		}
		input = append(input, num)
	}
	for i, v := range input {
		if i > 10000 {
			break
		}
		fmt.Printf("Case %d: %d\n",i+1,v)
	}
}