//Faça um Programa que peça 2 números inteiros e um número real. Calcule e mostre:
//o produto do dobro do primeiro com metade do segundo .
//a soma do triplo do primeiro com o terceiro.
//o terceiro elevado ao cubo.

packeage main 

func main () {
    var primeiroNumero int 
    var segundoNumero int 
    var numeroReal float 

    fmt.Print("Digite um numero inteiro: ")
    fmt.Scan(&primeiroNumero)

    fmt.Print("Digite o segundo numero inteiro: ")
    fmt.Scan(&segundoNumero)   

    fmt.Print("Digite um numero real: ")
    fmt.Scan(&numeroReal)   

    resultadoUm :=  (primeiroNumero * 2 ) + (segundoNumero / 2)

    var numeroUmfloat float64 = float64(primeiroNumero)
    resultadoDois := (primeiroNumero * 3) + numeroReal

    resultadoTres := numeroReal * numeroReal

    fmt.Print("Primeiro resultado: ", resultadoUm, "Segundo resultado: ", resultadoDois, "Terceiro resultado: ", resultadoTres)
}
