package main

import "fmt"

func main(){
	var n int
	var taro, hanako string
	var score [2]int

	fmt.Scan(&n)
	for i:=0; i<n; i++{
		fmt.Scanf("%s %s", &taro, &hanako)
		score = calc(taro, hanako, score)
	}
	fmt.Printf("%d %d\n", score[0], score[1])
}

func calc(taro, hanako string, score [2]int)[2]int{
	if taro == hanako{
		score[0] +=1 
		score[1] +=1 
	} else if taro < hanako {
		score[1] += 3
	} else{
		score[0] += 3
	}
	return score
}