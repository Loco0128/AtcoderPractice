package main

import "fmt"

func main(){
	var n, m, l int
	fmt.Scanf("%d %d %d", &n, &m, &l)
	matrix_a := make([][]int, n)
	matrix_b := make([][]int, m)
	output := make([][]int, n)

	for i := range matrix_a {
		matrix_a[i] = make([]int, m)
	}

	for i := range matrix_b {
		matrix_b[i] = make([]int, l)
	}

	for i := range output {
		output[i] = make([]int, l)
	}

	// 要素の入力
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix_a[i][j])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			fmt.Scan(&matrix_b[i][j])
		}
	}

	for i:=0; i<n;i++{
		for j:=0;j<l;j++{
			for k:=0; k<m;k++{
				output[i][j] += matrix_a[i][k] * matrix_b[k][j]
			}
		}
	}

	for _,row := range output {
		for i,col:= range row{
			fmt.Printf("%d",col)
			if i != l-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}