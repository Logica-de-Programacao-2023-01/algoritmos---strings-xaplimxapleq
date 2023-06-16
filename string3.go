package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	fmt.Print("digite uma frase:")
	fmt.Scan(&frase)

	var chartoreplace, replacementchar string
	fmt.Print("digite a letra para substituição:")
	fmt.Scan(&chartoreplace)
	fmt.Print("digite a letra que vai ser substituida:")
	fmt.Scan(&replacementchar)

	newString := strings.ReplaceAll(frase, chartoreplace, replacementchar)
	fmt.Println("resultado:", newString)

}
