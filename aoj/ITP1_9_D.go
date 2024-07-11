package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	var str string
	var orderNum int
	var commands []string
	fmt.Scan(&str)
	fmt.Scan(&orderNum)
	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<orderNum;i++{
		line,_ := reader.ReadString('\n')
		commands = append(commands,line)
	}

	for _,command := range commands {
		commandSlice := strings.Fields(command)
		a,_ := strconv.Atoi(commandSlice[1])
		b,_ := strconv.Atoi(commandSlice[2])
		switch commandSlice[0]{
		case "replace":
			p := commandSlice[3]
			str = str[:a]+p+str[b+1:]
		case "reverse":
			rev := str[a:b+1]
			str = str[:a]+reverseString(rev)+str[b+1:]
		case "print":
			fmt.Printf("%s\n",str[a:b+1])
		}
	}
}

// 文字列を逆順にする関数
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}