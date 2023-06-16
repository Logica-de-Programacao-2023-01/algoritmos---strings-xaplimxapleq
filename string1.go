package main

import (
	"fmt"
	"strings"
)

func main() {
	var Frase string
	fmt.Print("digite uma frase:")
	fmt.Scanln(&Frase)

	uppercase := strings.ToUpper(Frase)
	fmt.Println("resultado:", uppercase)
}
