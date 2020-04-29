package main

import (
	"fmt"
)

func main() {
	var S string
	K := "keyence"
	fmt.Scan(&S)
	for i := 0; i < len(K)-1; i++ {
		if K[i] != S[i] {
			if i < len(S) && S[i+1] == K[i] {
				i++
			}
			if S[len(S)-len(K[i+1:])-1:] == K[i:] {
				fmt.Println("YES")
				return
			} else {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
