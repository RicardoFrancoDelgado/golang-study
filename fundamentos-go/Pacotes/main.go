package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo com a função main")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}


/*
	Usar o ultimo nome depois da / para referenciar o pacote a ser usado
	ex: checkmail.funcaoX()

*/