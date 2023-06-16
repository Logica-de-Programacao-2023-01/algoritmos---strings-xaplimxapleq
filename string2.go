package main

import (
	"fmt"
	"strings"
)

func main() {
	var Frase string
	fmt.Print("digite uma frase:")
	fmt.Scanln(&Frase)

	noSpaces := strings.ReplaceAll(Frase, " ", "")
	fmt.Println("resultado:", noSpaces)
}
