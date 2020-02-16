package main

import "fmt"

func main() {
	var a, b, c int
	kawaisou := 0
	fmt.Scan(&a, &b, &c)
	if a == b {
		kawaisou++
	}
	if b == c {
		kawaisou++
	}
	if c == a {
		kawaisou++
	}
	if kawaisou == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
