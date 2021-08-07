package main

import "fmt"

func main() {
	lista := []string{"queijo", "iogurte", "maçã", "macarrão", "manga", "cebola"}
	tamanho := len(lista)
	fmt.Printf("\nA lista de compras possui %d itens:\n\n", tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("%d) %s\n", i+1, lista[i])
	}
}

