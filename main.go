package main

import (
	"fmt"
	"strconv"
	"strings"
)

// verificador de cpf
func validateCpf(a, b, c, d int) int {
	if a == c && b == d {
		fmt.Println("CPF valido! :)")
		return 1
	} else {
		fmt.Println("CPF invalido! :(")
		return 0
	}
	return 0
}

func CpfArray(cpf string) {
	numeros := []int{}

	//cpfLastDigits := strings.Split(cpf, "-")[1]
	cpfArrayFinal := strings.Split(cpf, "-")[1]
	cpfArrayFinalnum := strings.Split(cpfArrayFinal, "")

	cpfArray := strings.ReplaceAll(cpf, ".", "")
	cpfArray = strings.Split(cpfArray, "-")[0]

	for i := 0; i < len(cpfArray); i++ {
		char := string(cpfArray[i])
		num, _ := strconv.Atoi(char)
		numeros = append(numeros, num)
	}

	//calculo primeiro numero apos a -
	resultadosCalculados := [9]int{}
	for i := 10; i >= 2; {
		for j := 0; j <= 8; j++ {
			resultado := numeros[j] * i
			resultadosCalculados[j] = resultado
			i--
		}
	}

	//somando valor da multiplicação do array
	var resultadosCalculadosSomados int
	for i := 0; i < len(resultadosCalculados); i++ {
		resultadosCalculadosSomados = resultadosCalculadosSomados + resultadosCalculados[i]
	}

	//dividindo resultado da soma por 11
	var restoDaDivisao int = resultadosCalculadosSomados % 11
	var primeiroDigito int
	if restoDaDivisao < 2 {
		primeiroDigito = 0
	} else if restoDaDivisao > 2 {
		primeiroDigito = 11 - restoDaDivisao
	}

	//calculo segundo numero apos a -
	numeros = append(numeros, primeiroDigito)
	resultadosCalculados2 := [10]int{}
	for i := 11; i >= 2; {
		for j := 0; j <= 9; j++ {
			resultado := numeros[j] * i
			resultadosCalculados2[j] = resultado
			i--
		}
	}

	var resultadosCalculados2Somados int
	for i := 0; i < len(numeros); i++ {
		resultadosCalculados2Somados = resultadosCalculados2Somados + resultadosCalculados2[i]
	}
	segundoDigito := resultadosCalculados2Somados % 11
	if segundoDigito < 2 {
		segundoDigito = 0
	} else if segundoDigito > 2 {
		segundoDigito = 11 - segundoDigito
	}

	primeiroDigitoArray, err := strconv.Atoi(cpfArrayFinalnum[0])
	if err != nil {
		fmt.Println(err)
	}

	segundoDigitoArray, err := strconv.Atoi(cpfArrayFinalnum[1])
	if err != nil {
		fmt.Println(err)
	}
	validateCpf(primeiroDigitoArray, segundoDigitoArray, primeiroDigito, segundoDigito)

}

func main() {
	CpfArray("SEU_CPF_AQUI")
}
