package main

import "fmt"

func main(){
	var n int
  fmt.Scan(&n)

  input := make([]int, n)
  
  for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
  }

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d",input[i])
		if i != 0 {
			fmt.Print(" ")
		}
  }
	fmt.Printf("\n")
}
