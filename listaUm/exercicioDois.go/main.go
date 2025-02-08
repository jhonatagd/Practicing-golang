//Faça um Programa que peça um número e então 
// mostre a mensagem O número informado foi [número].
package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)
	fmt.Printf("O número informado foi %d\n", numero)
}