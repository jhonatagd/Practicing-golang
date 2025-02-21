//Tendo como dado de entrada a altura (h) de uma pessoa, 
// construa um algoritmo que calcule seu peso ideal, utilizando as seguintes fórmulas:
//Para homens: (72.7*h) - 58
//Para mulheres: (62.1*h) - 44.7

package main

func main() {
	var h float

	fmt.Println("Digite sua altura: ")
	fmt.Scan(&h)




	var sexo string
	fmt.Println("Digite se é homem ou mulher: ")
	fmt.Scan(&sexo)

	if sexo == "homem" {
		pHomem := (72.7*h) - 58
		fmt.Print("Seu peso ideal é de: ", pHomem)
		break
	} else if sexo == "mulher" {
		pMulher := (62.1*h) - 44.7
		fmt.Print("Seu peso ideal é de: ", pMulher)
		break
	} else {
		fmt.Print("Resposta invalida tente novamente")
	}

}