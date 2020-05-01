package main

import "fmt"

func main() {
	var K, S, ans int
	fmt.Scan(&K, &S)
	for x := 0; x <= K; x++ {
		for y := 0; y <= K; y++ {
			z := S - x - y
			if 0 <= z && z <= K {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
