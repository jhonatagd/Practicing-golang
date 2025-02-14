//Faça um Programa que pergunte quanto você ganha por hora e o número 
// de horas trabalhadas no mês. Calcule e mostre o total do seu salário no referido mês.
package main 

func main() {
	var ganhoHoras float
	var horasTrabalhadas float

	fmt.Print("Quanto voce ganha por hora: ")
	fmt.Scan(&ganhoHoras)

	fmt.Print("Quanto horas voce trabalha por mes: ")
	fmt.Scan(&horasTrabalhadas)

	calculaSalario := ganhoHoras * horasTrabalhadas

	fmt.Print("Seu salario será de: ",calculaSalario )
}