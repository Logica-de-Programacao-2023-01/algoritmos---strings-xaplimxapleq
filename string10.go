package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var fra1, fra2 string
	fmt.Print("Digite a primeira frase: ")
	fmt.Scanln(&fra1)
	fmt.Print("Digite a segunda frase: ")
	fmt.Scanln(&fra2)

	if isAnagram(fra1, fra2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

func isAnagram(fra1, fra2 string) bool {
	fra1 = strings.ToLower(fra1)
	fra2 = strings.ToLower(fra2)

	fra1 = sortString(fra1)
	fra2 = sortString(fra2)

	return fra1 == fra2
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
