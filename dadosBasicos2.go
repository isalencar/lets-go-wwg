package main

import (
	"fmt"
	"time"
)

func main() {
	var anoAtual int
	var anoNasci int
	var idade int
	anoAtual = time.Now().Year()
	fmt.Println("Quando você nasceu?")
	fmt.Scan(&anoNasci)
	idade = anoAtual - anoNasci

	fmt.Printf("Você tem %d anos", idade)
}
