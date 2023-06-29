package main

import (
	"fmt"
	"strconv"
)

func main() {
	var palavra string
	fmt.Println("Digite uma palavra:")
	fmt.Scan(&palavra)

	if isCrescente(palavra) {
		fmt.Println("A palavra é uma sequência numérica crescente.")
	} else {
		fmt.Println("A palavra não é uma sequência numérica crescente.")
	}
}

func isCrescente(palavra string) bool {
	for i := range palavra {
		if i > 0 {
			num1, err1 := strconv.Atoi(string(palavra[i-1]))
			num2, err2 := strconv.Atoi(string(palavra[i]))

			if err1 != nil || err2 != nil || num1 >= num2 {
				return false
			}
		}
	}
	return true
}
