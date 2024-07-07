package main

import "fmt"

func main(){
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	// M * N のスライスを作成
	matrix_a := make([][]int, n)
	matrix_b := make([]int, m)
	answer := make([]int,n)
	for i := range matrix_a {
		matrix_a[i] = make([]int, m)
	}
	// 要素の入力
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix_a[i][j])
		}
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&matrix_b[i])
	}

	for i,col := range matrix_a{
		for j,row := range col {
			answer[i] += row * matrix_b[j]
		}
	}

	for _,v := range answer{
		fmt.Printf("%d\n",v)
	}
}