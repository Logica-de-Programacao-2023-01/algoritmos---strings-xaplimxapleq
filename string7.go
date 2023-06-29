package main

import (
	"fmt"
	"regexp"
)

func main() {
	var frase string
	fmt.Print("escreva uma frase:")
	fmt.Scanln(&frase)
	if containsNumber(frase) {
		fmt.Println("A frase contém pelo menos um número.")
	} else {
		fmt.Println("A frase não tem números.")
	}
}

func containsNumber(frase string) bool {
	match, _ := regexp.MatchString("[0-9]", frase)
	return match
}
