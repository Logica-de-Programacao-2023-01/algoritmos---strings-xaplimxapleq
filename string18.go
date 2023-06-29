package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	if containsOnlyNumbers(str) {
		fmt.Println("A string contém apenas números.")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}

func containsOnlyNumbers(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
