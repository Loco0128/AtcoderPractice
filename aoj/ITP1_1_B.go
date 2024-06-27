package main

import "fmt"

func main () {
	var num int
	fmt.Scanf("%d",&num)

	result := num * num * num
	fmt.Printf("%d\n", result)  // 計算結果を表示
}