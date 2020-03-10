package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	low8 := int(float64(A) / 0.08)
	high8 := int((float64(A) + 0.9) / 0.08)
	low10 := int(float64(B) / 0.1)
	high10 := int((float64(B) + 0.9) / 0.1)
	if high8 < low10 || high10 < low8 {
		fmt.Println(-1)
	} else if low8 < low10 {
		fmt.Println(low10)
	} else {
		fmt.Println(low8)
	}
	fmt.Println(low8, high8, low10, high10)
}
