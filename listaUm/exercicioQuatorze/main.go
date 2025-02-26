// João precisa que você faça um programa que leia a variável peso (peso de peixes) e calcule o excesso. 
// Gravar na variável excesso a quantidade de quilos além do limite e na variável multa o
//  valor da multa que João deverá pagar. Imprima os dados do programa com as mensagens adequadas.

package main

import "fmt"

func main() {
    // Define o limite de peso
    const limite = 50
    const multaPorQuilo = 4.00

    // Lê o peso de peixes
    var peso float64
    fmt.Print("Digite o peso de peixes (em quilos): ")
    fmt.Scan(&peso)

    // Calcula o excesso, se houver
    var excesso float64
    var multa float64

    if peso > limite {
        excesso = peso - limite
        multa = excesso * multaPorQuilo
    }

    // Imprime o resultado
    if excesso > 0 {
        fmt.Printf("Excesso de peso: %.2f quilos\n", excesso)
        fmt.Printf("Valor da multa a pagar: R$ %.2f\n", multa)
    } else {
        fmt.Println("Não há excesso de peso, portanto não há multa.")
    }
}