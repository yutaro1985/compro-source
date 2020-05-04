package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)
	for i := 0; i < len(N); i++ {
		if N[i] != N[0] {
			fmt.Println("DIFFERENT")
			return
		}
	}
	fmt.Println("SAME")
}
