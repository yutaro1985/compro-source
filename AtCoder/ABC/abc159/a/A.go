package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	odds := N * (N - 1) / 2
	even := M * (M - 1) / 2
	fmt.Println(odds + even)
}
