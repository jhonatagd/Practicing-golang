//Faça um Programa que peça dois números e imprima a soma.
package main 

import (
	"fmt"
)

func main() {
	var numeroUm int
	var numeroDois int
	fmt.Print("Digite um numero:")
	fmt.Scan(&numeroUm)
	fmt.Print("Digite um numero:")
	fmt.Scan(&numeroDois)
	fmt.Printf("A soma dos numeros é de: %d\n", numeroUm + numeroDois)
}