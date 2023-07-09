package main

import "fmt"

func main() {
	fmt.Println("Write a progrram thath shows numbers from 1 to 100 but for multiples of 3 showss Shazam intead of the number and Shalakar for multiple of 5 and when is for both shows shakalakaboom")

	for i := 1; i <= 100; i++ {
		txt := ""
		if i % 3 == 0 {
			txt += "Fixx"
		}

		if i % 3 == 0 {
			txt += "buzz"
		}

		
//		fmt.Println(i)
//		var a int
//		a := i % 3
//		switch i {
//		case a == 0:
//			fmt.Println("Shakalaka")

		}
	}

}
