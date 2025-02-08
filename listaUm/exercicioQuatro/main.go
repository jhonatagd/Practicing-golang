//Faça um Programa que peça as 4 notas bimestrais e mostre a média.
package main

func main() {
	var notaUmBi float
	var notaDoisBi float
	var notaTresBi float
	var notaQuatroBi float

	fmt.Print("Digite a nota do primeiro semestre: ")
	fmt.Scan(&notaUmBi)
	fmt.Print("Digite a nota do segundo semestre: ")
	fmt.Scan(&notaDoisBi)
	fmt.Print("Digite a nota do terceiro semestre: ")
	fmt.Scan(&notaTresBi)
	fmt.Print("Digite a nota do quatro semestre: ")
	fmt.Scan(&notaQuatroBi)

	somaNota := notaUmBi + notaDoisBi + notaTresBi + notaQuatroBi / 4

	fmt.Printf("A media das notas são de: ", somaNota )
}