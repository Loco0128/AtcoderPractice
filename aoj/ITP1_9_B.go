package main

import "fmt"

type Data struct {
	dataset string
	h []int
}

func main(){
	var dataset string
	var num ,in int
	var data []Data
	var output []string
	for {
		var h []int
		fmt.Scan(&dataset)
		if dataset=="-"{
			break
		}
		fmt.Scan(&num)
		for i:=0;i<num;i++{
			fmt.Scan(&in)
			h = append(h, in)
		}
		data = append(data, Data{dataset,h})
	}

	for _,v := range data {
		after_data := shuffule(v)
		output = append(output,after_data)
	}
	for _,v := range output {
		fmt.Printf("%s\n",v)
	}
}

func shuffule(data Data) string{
	dataset := data.dataset
	for _,v := range data.h {
		dataset = dataset[v:] + dataset[:v]
	}

	return dataset
}