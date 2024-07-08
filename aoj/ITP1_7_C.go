package main

import "fmt"

func main(){
	var r, c int

	fmt.Scanf("%d %d", &r, &c)
	input := make([][]int, r + 1)
	for i := range input {
		input[i] = make([]int, c + 1)
	}

	// 要素の入力
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&input[i][j])
		}
	}
	// rowの合計を出す
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			input[i][c] += input[i][j]
		}
	}

	// calの合計を出す
	for i := 0; i <= c; i++ {
		for j := 0; j < r; j++ {
			input[r][i] += input[j][i]
		}
	}

	for _,row := range input {
		for i,col:= range row{
			fmt.Printf("%d",col)
			if i != c {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}