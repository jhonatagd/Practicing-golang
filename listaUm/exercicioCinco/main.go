//Faça um Programa que converta metros para centímetros.
package main

func(){
	var metros

	fmt.Print("Digite metros para converter em centímetros: ")
	fmt.Scan(&metros)
	convertMetros := metros * 100

	fmt.Print("Seu valor em centimetros: ", convertMetros)
}