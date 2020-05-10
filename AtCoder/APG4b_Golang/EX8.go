package main

import "fmt"

func main() {
	var P, price, N int
	fmt.Scan(&P)
	if P == 1 {
		fmt.Scan(&price, &N)
		fmt.Println(price * N)
	} else {
		var text string
		fmt.Scan(&text, &price, &N)
		fmt.Println(text + "!")
		fmt.Println(price * N)
	}
}
