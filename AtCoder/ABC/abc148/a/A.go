package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	check := map[int]bool{
		1: false,
		2: false,
		3: false,
	}
	check[A] = true
	check[B] = true
	for k, v := range check {
		if !v {
			fmt.Println(k)
			return
		}
	}
}
