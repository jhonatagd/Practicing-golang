//Tendo como dados de entrada a altura de uma pessoa, construa um algoritmo 
// que calcule seu peso ideal, usando a seguinte fórmula: (72.7*altura) - 58

package main

func main() {
	var alturaPessoa float

	fmt.Print("Digite sua altura: ")
	fmt.Scan(&alturaPessoa)

	calculaPeso := (72.7 * alturaPessoa) - 58

	fmt.Print("Seu peso ideal: ", calculaPeso)
}