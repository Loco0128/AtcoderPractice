package main

import "fmt"

func main(){
	var input [][]int
	var numa, numb, tmp int

	for {
		fmt.Scanf("%d %d", &numa,&numb)
		if numa == 0 && numb ==0 {
			break
		}
		if numa > numb {
			tmp = numa
			numa = numb
			numb = tmp
		}
		input_data := []int{numa, numb}
		input = append(input, input_data)
	}
	for _, v := range input {
		fmt.Printf("%d %d\n",v[0],v[1])
	}
}