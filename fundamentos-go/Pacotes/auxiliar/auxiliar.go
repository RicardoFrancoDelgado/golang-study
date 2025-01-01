package auxiliar

import "fmt"

// Escrever registra no console uma mensagem
func Escrever() {
	fmt.Println("Escrevendo com a função do pacote auxiliar")
	escrever2()
}

/*
	Em Go precisamos fazer com que a função seja iniciada com letra maiúscula para exportar
	a mesma em outro arquivo, já que ela vem de outro pacote.
	No mesmo pacote podemos utilizar funções com letras minúsculas, elas serão executadas dentro da função com nome maiúsculo
	e isso fará com que ela não seja chamada em um arquivo main por ex.
*/