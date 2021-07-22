package main

import (
	"fmt"
)

func main() {
	//Não compila porque o tipo 'int8' não contém bits o bastante.
	//O erro nos indica que o número 150 não cabe no espaço reservado.

	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}
