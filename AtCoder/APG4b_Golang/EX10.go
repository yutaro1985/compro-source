package main

import "fmt"

func main() {
	// goにはWhile文がない
	var A, B int
	fmt.Scan(&A, &B)
	Ascore := "A:"
	Bscore := "B:"
	for i := 0; i < A; i++ {
		Ascore += "]"
	}
	for i := 0; i < B; i++ {
		Bscore += "]"
	}
	fmt.Println(Ascore)
	fmt.Println(Bscore)
}
