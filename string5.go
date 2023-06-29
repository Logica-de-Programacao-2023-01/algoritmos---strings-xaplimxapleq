package main

import (
	"fmt"
	"strconv"
)

func main() {
	var caractere string
	fmt.Print("Digite um caractere em ponto flutuante: ")
	fmt.Scanln(&caractere)

	_, err := strconv.ParseFloat(caractere, 64)
	if err == nil {
		fmt.Println("É um caractere válido em ponto flutuante.")
	} else {
		fmt.Println("Não é um caractere válido em ponto flutuante.")
	}
}
