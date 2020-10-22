package main

import "fmt"

func main() {
	var h, a int
	fmt.Scan(&h)
	fmt.Scan(&a)
	if h%a == 0 {
		fmt.Println(h / a)
	} else {
		fmt.Println(h/a + 1)
	}
}
