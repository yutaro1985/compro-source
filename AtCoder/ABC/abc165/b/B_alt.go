package main

import "fmt"

func main() {
	var X, ans int
	fmt.Scan(&X)
	for i := 100; i < X; i += i / 100 {
		ans++
	}
	fmt.Println(ans)
}
