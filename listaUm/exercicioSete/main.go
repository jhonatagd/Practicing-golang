//Faça um Programa que calcule a área de um quadrado, 
// em seguida mostre o dobro desta área para o usuário.
package main() {
	var areaQuadrado float

	fmt.print("Digite o numero para calcular a area de um quadrado: ")
	fmt.Scan(&areaQuadrado)

	formulaQuadrada := areaQuadrado * areaQuadrado

	fmt.Print("Area calculada: ", formulaQuadrada)
}