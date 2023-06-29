package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&palavra)

	result := removeVowels(palavra)
	fmt.Println("Resultado:", result)
}

func removeVowels(str string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}
