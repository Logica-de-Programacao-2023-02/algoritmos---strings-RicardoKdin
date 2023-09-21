package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	var x string
	fmt.Print("Digite uma string: ")
	fmt.Scan(&x)

	isCamelCase, wordCount := isCamelCase(x)

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavra(s).\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func isCamelCase(s string) (bool, int) {

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsUpper(r)
	})

	for _, word := range words {
		if !isCamelCaseWord(word) {
			return false, len(words)
		}
	}

	return true, len(words)
}

func isCamelCaseWord(word string) bool {

	if len(word) < 1 {
		return false
	}

	firstChar := rune(word[0])
	if !unicode.IsUpper(firstChar) {
		return false
	}

	for _, char := range word[1:] {
		if !unicode.IsLower(char) {
			return false
		}
	}

	return true
}
