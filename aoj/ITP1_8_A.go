package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// 文字列をルーンのスライスに変換
	runes := []rune(input)

	for i, char := range runes {
		if unicode.IsUpper(char) {
			runes[i] = unicode.ToLower(char)
		} else if unicode.IsLower(char) {
			runes[i] = unicode.ToUpper(char)
		}
	}

	// ルーンのスライスを文字列に戻す
	output := string(runes)

	fmt.Print(output)
}
