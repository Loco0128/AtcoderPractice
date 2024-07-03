package main

import (
    "fmt"
		"math"
)

func main() {
  var n int
  fmt.Scan(&n)

  input := make([]int, n)
  
  for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
  }

	min := min(input)
	max := max(input)
	sum := sum(input)
	fmt.Printf("%d %d %d\n",min, max, sum)
}

func min(input []int) int {
	var min = math.MaxInt64
	for _,v := range input {
		if v < min {
			min = v
		}
	}
	return min
}

func max(input []int) int {
	var max = math.MinInt64
	for _,v := range input {
		if v > max {
			max = v
		}
	}
	return max
}

func sum(input []int) int {
	var sum = 0
	for _,v := range input {
		sum = v + sum
	}
	return sum
}