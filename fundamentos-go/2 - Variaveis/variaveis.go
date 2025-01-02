package main

import "fmt"

func main() {
	// declarar variáveis de forma explicita
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	//declarar variáveis de forma implicita com inferência de tipo
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	//declarar multiplas variáveis encadeadas
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3, variavel4)

	//declarar multiplas variaveis encadeadas utilizando inferência de tipo
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	//invertendo o valor de uma variável com a outra sem a necessidade de criar uma variável auxiliar
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}