package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(cpf string) {
	// Separar os dígitos do CPF
	parteNumerica := strings.ReplaceAll(cpf, ".", "")
	partes := strings.Split(parteNumerica, "-")
	numerosStr := partes[0]
	digitosFinais := strings.Split(partes[1], "")

	// Converter para array de inteiros
	numeros := []int{}
	for _, char := range numerosStr {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Erro ao converter CPF para inteiro:", err)
			return
		}
		numeros = append(numeros, num)
	}

	// Calcular primeiro dígito
	primeiro := calcularDigito(numeros, 10)
	digito1, _ := strconv.Atoi(digitosFinais[0])

	// Calcular segundo dígito
	numeros = append(numeros, primeiro)
	segundo := calcularDigito(numeros, 11)
	digito2, _ := strconv.Atoi(digitosFinais[1])

	CpfVerify(primeiro, segundo, digito1, digito2)
}

func calcularDigito(numeros []int, pesoInicial int) int {
	soma := 0
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i] * (pesoInicial - i)
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func CpfVerify(calculado1, calculado2, informado1, informado2 int) {
	if calculado1 == informado1 && calculado2 == informado2 {
		fmt.Println("CPF válido! :)")
	} else {
		fmt.Println("CPF inválido! :(")
	}
}
