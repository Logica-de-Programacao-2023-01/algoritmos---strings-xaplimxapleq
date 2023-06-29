package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Digite uma frase: ")
	var str string
	fmt.Scan(&str)

	fmt.Print("Digite a letra a ser substituída: ")
	var oldChar string
	fmt.Scan(&oldChar)

	fmt.Print("Digite a letra substituta: ")
	var newChar string
	fmt.Scan(&newChar)

	replaced := strings.Map(func(r rune) rune {
		if string(r) == oldChar {
			return []rune(newChar)[0]
		}
		return r
	}, str)

	fmt.Printf("A string \"%s\" com \"%s\" substituído por \"%s\" é \"%s\".\n", str, oldChar, newChar, replaced)
}
