//Faça um Programa que peça a temperatura em graus Celsius,
//  transforme e mostre em graus Fahrenheit.
package main 

func main() {
	var temperaturaC float

	fmt.Print("Digite uma temperatura em graus Celsius: ")
	fmt.Scan(&temperaturaC)

	convertTemperatura := 5 / 9 * (temperaturaC - 32)

	fmt.Print("Sua temperatura em graus Fahrenheit: ", convertTemperatura)
}