package main

import (
	"bufio"
	"fmt"
	"sort"
	"os"
	"unicode"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var inputLines []string
	for {
		// 1行読み込む
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// 入力を保存
		inputLines = append(inputLines, line)
	}

	// すべての行を結合して1つの文字列にします
	input := strings.Join(inputLines, "")

	lettersMap := make(map[rune]int)
  for r := 'a'; r <= 'z'; r++ {
  	lettersMap[r] = 0
  }

	// 文字列をルーンのスライスに変換
	runes := []rune(input)

	for _, char := range runes {
		if unicode.IsLetter(char) {
			lettersMap[unicode.ToLower(char)] += 1
		}
	}


	var letters []rune
	for letter := range lettersMap {
		letters = append(letters, letter)
	}
	
	// スライスをアルファベット順にソートします
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	
	// ソート後のアルファベットとその対応する値を表示します
	for _, letter := range letters {
		fmt.Printf("%c : %d\n", letter, lettersMap[letter])
	}
}
