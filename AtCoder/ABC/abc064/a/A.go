package main

import "fmt"

func main() {
	var r, g, b int
	fmt.Scan(&r, &g, &b)
	if (100*r+10*g+b)%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
