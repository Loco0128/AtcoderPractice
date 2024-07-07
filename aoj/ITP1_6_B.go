package main

import "fmt"

type CardType struct {
	kind string
	num int
}

func creatAllCard()[]CardType{
	var card_list []CardType
	var cardKind = []string{"S","H","C","D"}
	for _,v := range cardKind {
		for j:=0;j<13;j++{
			card_list = append(card_list, CardType{v , j+1})
		}
	}
	return card_list
}

func removeIndex(s []CardType, index int) []CardType {
	return append(s[:index], s[index+1:]...)
}

func removeCard(card_list []CardType, value CardType)[]CardType{
	for i,v := range card_list{
		if v.kind == value.kind && v.num == value.num{
			return removeIndex(card_list,i)
		}
	}
	return card_list
}

func main(){
	var n, num int
	var kind string

	output := creatAllCard()
	fmt.Scan(&n)
  
  for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &kind, &num)
		output = removeCard(output, CardType{kind, num})
  }

	for _,v := range output {
		fmt.Printf("%s %d\n", v.kind, v.num )
	}
}