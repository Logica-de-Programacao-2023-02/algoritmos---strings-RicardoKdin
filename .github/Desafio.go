package main

import (
	"fmt"
	"strings"
)

func main() {

	var x, padrao string

	fmt.Print("Escreva uma string:")
	fmt.Scan(&x)
	fmt.Println("Escreva um padrão: ")
	fmt.Scan(&padrao)

	var index []int

	for i := len(x) - 1; i >= 0; i -= len(padrao) {
		idx := strings.Index(x[i:], padrao)
		if idx != -1 {
			index = append(index, i+idx)
		}
		fmt.Println(index)
	}
}
