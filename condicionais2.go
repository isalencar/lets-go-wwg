package main

import "fmt"

func main() {
	var numero int
	fmt.Scan(&numero)
	if numero == 0{
		fmt.Printf("este numero é 0")
	}else if numero % 2 == 0{
		if numero % 3 == 0{
			fmt.Printf("este numero é divisível por 2 e por 3")
			return
		}
		fmt.Printf("este numero é divisível por 2")
	}else if numero % 3 == 0{
		fmt.Printf("este numero é divisível por 3")
	}else {
		fmt.Printf("este numero não é 0, nem divisível por 2 ou por 3")
	}

}