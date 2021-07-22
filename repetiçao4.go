package main

import "fmt"

func main() {
	var i, numero int
	for i == 0{
		fmt.Printf("Insira um número:\n")
		fmt.Scan(&numero)
		if numero % 2 == 0{
			i++
		}
	}
	fmt.Printf("Agora sim podemos dividir igualmente entre mim e você!\n")
}
