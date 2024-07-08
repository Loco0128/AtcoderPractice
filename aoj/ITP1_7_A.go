package main

import "fmt"

func main(){
	var m,f,r int
	var input [][3]int
	for i:=0;i<50;i++{
		fmt.Scanf("%d %d %d", &m, &f, &r)
		if m == -1 && f == -1 && r == -1{
			break
		}
		input = append(input, [3]int{m,f,r})
	}
	output := grade(input)

	for _,v := range output{
		fmt.Printf("%s\n",v)
	}
}

func grade(input [][3]int) []string{
	var output []string
	for _,v:= range input {
		sum :=  v[0] + v[1]
		if v[0] == -1 || v[1] == -1{
			output = append(output, "F")
		} else if sum >= 80 {
			output = append(output, "A")
		} else if sum >= 65 {
			output = append(output, "B")
		} else if sum >= 50 {
			output = append(output, "C")
		} else if sum >= 30 {
			if v[2] >= 50 {
				output = append(output, "C")
			} else{
				output = append(output, "D")
			}
		} else {
			output = append(output, "F")
		}
	}
	return output
}