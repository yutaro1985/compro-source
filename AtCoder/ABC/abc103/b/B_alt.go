package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)
	for i := 0; i < len(S); i++ {
		if T == S[i:]+S[:i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
