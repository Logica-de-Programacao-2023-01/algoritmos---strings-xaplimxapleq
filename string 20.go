package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	isCamelCase, count := checkCamelCase(str)

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavras.\n", count)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func checkCamelCase(str string) (bool, int) {
	words := splitCamelCase(str)
	count := len(words)

	if count == 0 || !unicode.IsUpper(rune(str[0])) {
		return false, 0
	}

	return true, count
}

func splitCamelCase(str string) []string {
	splitFunc := func(r rune) bool {
		return unicode.IsUpper(r)
	}

	return strings.FieldsFunc(str, splitFunc)
}
