package main

import (
	"fmt"
	"math"
)

func main(){
	var n int
	var x, y float64
	var array_x, array_y []float64

	fmt.Scan(&n)
	for i:=0; i<n;i++{
		fmt.Scan(&x)
		array_x = append(array_x, x)
	}
	for i:=0; i<n;i++{
		fmt.Scan(&y)
		array_y = append(array_y, y)
	}

	fmt.Printf("%.6f\n",calcMh(array_x, array_y, n))
	fmt.Printf("%.6f\n",calcUc(array_x, array_y, n))
	fmt.Printf("%.6f\n",calcMin(array_x, array_y, n))
	fmt.Printf("%.6f\n",calcCb(array_x, array_y, n))
}

func calcMh (x,y []float64, leng int)float64{
	var ans float64
	for i:=0;i<leng;i++{
		ans += abs(x[i]-y[i])
	}
	return ans
}

func calcUc (x,y []float64, leng int)float64{
	var ans float64
	for i:=0;i<leng;i++{
		ans += math.Pow(abs(x[i]-y[i]),2)
	}
	return math.Sqrt(ans)
}

func calcMin (x,y []float64, leng int)float64{
	var ans float64
	for i:=0;i<leng;i++{
		ans += math.Pow(abs(x[i]-y[i]),3)
	}
	return math.Pow(ans, 1.0/3.0)
}

func calcCb (x,y []float64, leng int)float64{
	var max float64
	for i:=0;i<leng;i++{
		if abs(x[i]-y[i]) > max {
			max = abs(x[i]-y[i])
		}
	}
	return max
}

func abs(num float64)float64{
	if num > 0{
		return num
	} else {
		return -num
	}
}
