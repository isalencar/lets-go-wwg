package main

import "fmt"

func main() {
	ilhas := map[string]string{
		"Santorini" : "Grécia",
		"Ilha de Creta" : "Grécia",
		"Koh Phi Phi" : "Tailândia",
		"Koh Samui" : "Tailândia",
		"Ilhas dos Açores" : "Portugal",
		"Bora Bora" : "PolinésiaFrancesa",
		"Aruba" : "Caribe",
		"Bali" : "Indonésia",
		"Arquipélago das Antilhas" : "Caribe",
		"Bahamas" : "Caribe",
	}
	freqPaises := make(map[string]int)
	for _, i := range ilhas{
		freqPaises[i]++
	}
	fmt.Println(freqPaises)
}