package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos int, int8 - int16 - int32 - int64
	// declaração literal
	var inteiro int8 = 10
	var inteiro2 int16 = 100
	var inteiro3 int32 = 1000
	var inteiro4 int64 = 1000000
	var inteiro5 int = -100000000
	fmt.Println(inteiro, inteiro2, inteiro3, inteiro4, inteiro5)

	// declaração por inferência de tipo
	inteiro6 := 50
	fmt.Println(inteiro6)

	// unsigned int: uint, uint8 - uint16 - uint32 - uint64
	var positive7 uint = 100
	fmt.Println(positive7)
	//FIM DO INT/UINT

	//Tipo Float - float32 - float64
	var float1 float32 = 129.99
	var float2 float64 = 1777.88
	fmt.Println(float1, float2)

	//declaração por inferência de tipo - float
	float3 := 123.45
	fmt.Println(float3)
	//FIM DO FLOAT

	// Tipo String e "CHAR"
	var str string = "Texto"
	fmt.Println(str)


	//CHAR é tratado como int ou rune - int32
	char := 'B' // 66 na tabela ascii
	fmt.Println(char)
	
	
	//Saidas 0 => Valor padrão de variáveis que não são iniciadas
	var texto string
	var number int
	var erro error
	fmt.Println(texto) // string vazia
	fmt.Println(number) // 0
	fmt.Println(erro) // <nil>
	// FIM SAIDAS 0

	// Tipo Boolean
	var rain bool = false
	cold := true
	fmt.Println(rain, cold)
	// FIM BOOLEAN

	// Tipo error
	var validateError error = errors.New("Tudo errado")
	fmt.Println(validateError)
}