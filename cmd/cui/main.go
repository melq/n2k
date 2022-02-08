package main

import (
	"fmt"
	"n2k"
)

func main() {
	numStr := "123"
	kanji, err := n2k.Number2kanji(numStr)
	if err != nil {
		return
	}

	number, err := n2k.Kanji2number(kanji)
	if err != nil {
		return
	}
	fmt.Println(kanji, number)
}