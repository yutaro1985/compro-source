package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	num := 0
	for {
		if a%10 != 0 {
			num++
		} else if a/10 == 0 {
			break
		}
		a /= 10
	}
	fmt.Println(num)
}
