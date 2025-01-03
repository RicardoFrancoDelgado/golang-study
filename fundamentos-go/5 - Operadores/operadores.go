package main

import "fmt"

func main() {
	// Operadores aritméticos
	fmt.Println(1 + 2)// soma
	fmt.Println(1 - 2)// subtracao
	fmt.Println(1 / 2)// divisao
	fmt.Println(1 * 2)// multiplicacao
	fmt.Println(1 % 2)// mod / resto da divisao

	// Operadores de atribuição
	var texto string = "Texto" // atribuição explicita
	texto2 := "Texto 2" // atribuição por inferência
	fmt.Println(texto, texto2)

	// Operadores relacionais
	fmt.Println(1 > 2) // maior
	fmt.Println(1 >= 2) // maior igual
	fmt.Println(1 == 2) // igual de comparação
	fmt.Println(1 < 2) // menor
	fmt.Println(1 <= 2) // menor igual
	fmt.Println(1 != 2) // diferença

	// Operadores lógicos
	fmt.Println("----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // E -> And
	fmt.Println(verdadeiro || falso) // OU -> OR
	fmt.Println(!verdadeiro) // Negação -> NOT
	fmt.Println(!falso)

	// Operadores unários
	numero := 30
	// incrementar
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)
	numero *= 2
	fmt.Println(numero)
	// decrementar
	numero--
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
	numero /= 2
	fmt.Println(numero)
	numero %= 2
	fmt.Println(numero)

	// Operador ternário => Não existem operadores ternários em Go
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	// caso queira resolver de forma pareida use if else
	if numero > 5 {
		fmt.Println("Maior que 5")
	} else {
		fmt.Println("Menor que 5")
	}
}