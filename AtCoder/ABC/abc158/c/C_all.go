package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	// 1009までなので普通に全探索でもOK
	for i := 1; i <= 1009; i++ {
		if int(float64(i)*0.08) == A && int(float64(i)*0.1) == B {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
