package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()

	if containsNumber(frase) {
		fmt.Println("A frase contém um número")
	} else {
		fmt.Println("A frase não contém um número")
	}

}
func containsNumber(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
