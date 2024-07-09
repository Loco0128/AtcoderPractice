package main

import (
	"fmt"
	"strings"
)

func main(){
	var text, substring string
	fmt.Scan(&text)
	fmt.Scan(&substring)
	if strings.Contains(text+text, substring) {
		fmt.Print("Yes\n")
	} else {
		fmt.Print("No\n")
	}
}