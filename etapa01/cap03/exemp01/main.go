package main

import "fmt"

func main() {
	var x string
	x = "Hallo Welt!!"
	fmt.Println(x)

	x += " Segundo"

	fmt.Println(x)

	var y string = "Ciao Mundo!"
	var z string = "Ariveredeti"
	fmt.Println(y == z)
	var w string
	w = "Hallo Welt!!"
	fmt.Println(x == w)

	// should use camelCase
}
