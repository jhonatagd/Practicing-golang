//Faça um Programa que peça a temperatura em graus
// Fahrenheit, transforme e mostre a temperatura em graus Celsius
package main 

func main() {
	var temperaturaF float

	fmt.Print("Digite uma temperatura em graus Fahrenheit: ")
	fmt.Scan(&temperaturaF)

	convertTemperatura := 5 * ((temperaturaF-32)/ 9)

	fmt.Print("Sua temperatura em graus Celsius: ", convertTemperatura)
}