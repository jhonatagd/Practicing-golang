//Faça um Programa que peça o raio de um círculo
//calcule e mostre sua área.
package main

func main(){
	var raio float
	var formulaRaio

	fmt.Print("digite um raio de um circulo: ")
	fmt.Scan(&raio)
	formulaRaio = raio / 2.0

	fmt.Print("Valor do diametro: ", formulaRaio)
}