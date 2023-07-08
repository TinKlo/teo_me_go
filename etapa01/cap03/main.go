package main

import "fmt"

func main() {
	fmt.Println("Entre com o valor em metros: ")
	var p float64
	fmt.Scanf("%f", &p)

	m := p * 0.348
	fmt.Println(m, "metros")
}
