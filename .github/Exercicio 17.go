package main

import (
	"fmt"
)

func main() {

	var x string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&x)

	letras := unicas(x)
	fmt.Println("Letras únicas na palavra:", letras)
}

func unicas(s string) string {
	letras := ""
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	for _, char := range s {
		if counts[char] == 1 {
			letras += string(char)
		}
	}

	return letras
}
