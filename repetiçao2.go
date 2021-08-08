package main

import "fmt"

type Apartamento struct{
	dono string
	númeroApto int
	temVaga bool
}
func main() {
	apartamento1 := Apartamento{
		dono: "Luisa",
		númeroApto: 107,
		temVaga: true,
	}
	apartamento2 := Apartamento{
		dono: "Ana",
		númeroApto: 113,
		temVaga: false,
	}
	apartamento3 := Apartamento{
		dono: "Maria",
		númeroApto: 207,
		temVaga: true,
	}
	apês := []Apartamento{apartamento1, apartamento2, apartamento3}

	for _, apartamento := range apês {
		fmt.Printf("O apartamento da %s é o número %d e ", apartamento.dono, apartamento.númeroApto)
		if apartamento.temVaga {
			fmt.Printf("possui vaga de estacionamento.\n")
		}else{
			fmt.Printf("não possui vaga de estacionamento.\n")
		}
	}


}
