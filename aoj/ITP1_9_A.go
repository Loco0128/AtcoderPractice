package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main (){
	var w string
	var count int
	fmt.Scan(&w)
	w = strings.ToLower(w)

	reader := bufio.NewReader(os.Stdin)
	var inputLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line=="END_OF_TEXT\n" {
			break
		}
		inputLines = append(inputLines, line)
	}
	
	for _, line := range inputLines {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word) 
			if word == w {
				count++
			}
		}
	}

	fmt.Println(count)
}