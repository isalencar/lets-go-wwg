package main

import "fmt"

func main() {
	var a, b, c, maior int
	fmt.Scan(&a, &b, &c)
	if a >= b && a >= c{
		maior = a
	}else if b >= a && b >= c{
		maior = b
	}else {
		maior = c
	}
	fmt.Println(maior)
}