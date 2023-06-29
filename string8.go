package main

import "fmt"

func main() {
	var frase string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&frase)

	reversed := reverseString(frase)
	fmt.Println("frase invertida:", reversed)
}

func reverseString(frase string) string {
	runes := []rune(frase)
	for e, o := 0, len(runes)-1; e < o; e, o = e+1, o-1 {
		runes[e], runes[o] = runes[o], runes[e]
	}
	return string(runes)
}
