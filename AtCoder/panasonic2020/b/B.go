package main

import "fmt"

func main() {
	var H, W int
	var odd, even int
	fmt.Scan(&H, &W)
	if H == 1 || W == 1 {
		fmt.Println(1)
		return
	}
	if W%2 != 0 {
		even = W/2 + 1
		odd = W / 2
	} else {
		even = W / 2
		odd = W / 2
	}
	if H%2 != 0 {
		fmt.Println((H/2+1)*even + H/2*odd)
	} else {
		fmt.Println(H/2*even + H/2*odd)
	}

}
