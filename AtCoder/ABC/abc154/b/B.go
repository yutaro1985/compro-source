package main

import "fmt"

func main() {
	var S, ans string
	fmt.Scan(&S)
	for i := 0; i < len(S); i++ {
		ans += "x"
	}
	fmt.Println(ans)
}
