package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)
	x++
	fmt.Println(x)
	x *= a + b
	fmt.Println(x)
	x *= x
	fmt.Println(x)
	x--
	fmt.Println(x)
}
