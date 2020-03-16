package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	list := make(map[int]int)
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)
		if _, exist := list[A]; exist {
			fmt.Println("NO")
			return
		} else {
			list[A] = A
		}
	}
	fmt.Println("YES")
}
