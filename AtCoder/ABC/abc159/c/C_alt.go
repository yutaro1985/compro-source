package main

import "fmt"

func main() {
	var L float64
	fmt.Scan(&L)
	a := L / 3.0
	b := L / 3.0
	c := L / 3.0
	ans := a * b * c
	fmt.Println(ans)
}
