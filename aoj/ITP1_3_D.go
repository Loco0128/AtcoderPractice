package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scanf("%d %d %d",&a, &b, &c)
	array := divisor(c)
	num := counter(a,b,array)
	fmt.Printf("%d\n",num)
}

func divisor (c int) []int{
	var num []int
	for i:=1; i<=c;i++{
		if c % i == 0{
			num = append(num, i)
		}
	}
	return num
}

func counter (from int, to int, array []int) int{
	var count int
	for i:=from; i<= to; i++{
		for _,v := range array {
			if v == i {
				count++
			}
		}
	} 
	return count
}