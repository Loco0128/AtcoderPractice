package main

import "fmt"

type Operation struct {
	a  int
	op string
	b  int
}

func main(){
	var a, b, result int
	var op string
	var input []Operation
	for {
		fmt.Scanf("%d %s %d", &a, &op, &b) 
		if op == "?"{
			break
		}
		input =  append(input,Operation{a,op,b})
	}

	for _,v := range input {
		switch v.op {
		case "+":
			result = v.a + v.b
		case "-":
			result = v.a - v.b
		case "*":
			result = v.a * v.b
		case "/":
			result = v.a / v.b
		}
		fmt.Printf("%d\n",result)
	}
}