package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Println("Digite uma palavra:")
	fmt.Scan(&palavra)

	palavra = replaceVogais(palavra, "*")

	fmt.Println(palavra)
}

func replaceVogais(palavra string, substituto string) string {
	vogais := "aeiouAEIOU"

	for _, vogal := range vogais {
		palavra = strings.ReplaceAll(palavra, string(vogal), substituto)
	}

	return palavra
}
