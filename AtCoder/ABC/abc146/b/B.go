package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	ans := make([]byte, len(S))
	for i := 0; i < len(S); i++ {
		ans[i] = (S[i]-'A'+byte(N))%26 + 'A'
	}
	fmt.Println(string(ans))
}
