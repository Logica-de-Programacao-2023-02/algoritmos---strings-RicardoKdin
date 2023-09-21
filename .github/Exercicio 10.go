package main

import "fmt"

func main() {

	var x, y string

	fmt.Print("Escreva duas palavras: ")
	fmt.Scan(&x, &y)

	xMapa := make(map[string]int)
	yMapa := make(map[string]int)

	for _, c := range x {
		xMapa[string(c)]++
	}
	for _, c := range y {
		yMapa[string(c)]++
	}

	fmt.Println(xMapa)
	fmt.Println(xMapa)

	for char, qtdX := range xMapa {
		qtdY := yMapa[char]
		if qtdX != qtdY {
			fmt.Println("As palavras não são anagramas")
			return
		} else {
			fmt.Println("As palavras são anagramas")
			return
		}
	}
}
