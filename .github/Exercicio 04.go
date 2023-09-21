package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	s1 := bufio.NewScanner(os.Stdin)
	s2 := bufio.NewScanner(os.Stdin)

	fmt.Print("Escreva uma frase: ")
	s1.Scan()
	f1 := s1.Text()

	fmt.Println("Escreva outra frase: ")
	s2.Scan()
	f2 := s2.Text()

	if f1 == f2 {
		fmt.Println("As frases são iguais")
	} else {
		fmt.Println("As frases são diferentes")
	}

}
