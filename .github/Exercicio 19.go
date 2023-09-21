package main

import "fmt"

func main() {

	var x string

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&x)

	var reverso string

	for i := len(x) - 1; i >= 0; i-- {
		reverso = reverso + string(x[i])
	}
	fmt.Println(reverso)
}
