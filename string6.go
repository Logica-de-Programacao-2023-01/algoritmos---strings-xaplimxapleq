package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("escreva uma frase:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	wordCount := countWords(input)
	fmt.Printf("A frase possui %d palavra(s).\n", wordCount)
}

func countWords(input string) int {
	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		return 0
	}

	words := strings.Fields(trimmedInput)
	return len(words)

}
