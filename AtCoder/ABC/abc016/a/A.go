package main

import "fmt"

func main() {
	var M, D int
	fmt.Scan(&M, &D)
	if M%D == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
