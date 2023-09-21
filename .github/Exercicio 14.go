package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x string

	fmt.Print("Escreva uma sequência: ")
	fmt.Scan(&x)

	ints := make([]int, len(x))
	for i, c := range x {
		var err error
		ints[i], err = strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("A sequência não é numérica")
			return
		}
	}
	isDecrescente := true
	for i := 1; i < len(ints); i++ {
		atual := ints[i]
		prox := ints[i-1]
		if atual <= prox {
			isDecrescente = false
			break
		}
	}
	if isDecrescente {
		fmt.Println("Não é decrescente")
	} else {
		fmt.Println("É decrescente")
	}
}
