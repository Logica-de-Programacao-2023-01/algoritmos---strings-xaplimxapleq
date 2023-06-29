package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	uniqueLetters := findUniqueLetters(str)

	fmt.Print("Letras únicas: ")
	for _, letter := range uniqueLetters {
		fmt.Printf("%c ", letter)
	}
	fmt.Println()
}

func findUniqueLetters(str string) []rune {
	count := make(map[rune]int)

	for _, char := range str {
		count[char]++
	}

	var uniqueLetters []rune
	for char, freq := range count {
		if freq == 1 {
			uniqueLetters = append(uniqueLetters, char)
		}
	}

	return uniqueLetters
}
