package main

import "fmt"

// declaração padrão de função com o especificação do tipo de retorno
func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//funções podem conter mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := soma(10, 20) // var soma do tipo func
	fmt.Println(soma)

	// atribuir a uma variável uma função com retorno e chamar a função a partir do nome da variável
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// informar que somente o primeiro valor será retornado ignorando o segundo com um underline
	somando, _ := calculosMatematicos(1, 10)
	fmt.Println(somando)
}