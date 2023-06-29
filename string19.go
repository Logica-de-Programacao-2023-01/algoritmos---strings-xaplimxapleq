package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	inverted := reverseString(str)

	fmt.Println(inverted)
}

func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		j := length - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
