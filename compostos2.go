package main

import "fmt"

func main() {
	amarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	vermelho := []string{"Helena", "Jonas", "José", "Juliana"}
	vermelho = append(vermelho, "Luis Inácio")
	fmt.Printf("amarelo: %v\nvermelho: %v\n", amarelo, vermelho)
}
