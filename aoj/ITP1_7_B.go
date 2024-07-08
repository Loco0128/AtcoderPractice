package main

import "fmt"

func main(){
	var n,x int
	var output []int

	for {
		fmt.Scanf("%d %d", &n, &x)
		if n == 0 && x == 0 {
			break
		}
		count:=calc(n,x)
		output = append(output,count)
	}

	for _,v := range output{
		fmt.Printf("%d\n",v)
	}
}

func calc(n, x int) int {
	var remain, count int
	for i:=1; i <= n; i++{
		for j:=i+1; j<=n;j++{
			remain = x - (i + j)
			if remain > 0 && remain <= n && remain > i && remain > j{
				count++
			}
		}  
	}
	return count
}