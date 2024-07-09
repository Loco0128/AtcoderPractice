package main

import (
	"fmt"
	"strconv"
)

func main(){
	var n string
	var input []string
	var output []int
	for {
		fmt.Scan(&n)
		if n == "0" {
			break
		}
		input = append(input, n)
	}
	for _,strNum := range input {
		var sum int
		for _,v:= range strNum{
			num ,_ := strconv.Atoi(string(v))
			 sum += num
		}
		output = append(output, sum)
	}
	for _,row := range output {
		fmt.Printf("%d\n",row)
	}
}