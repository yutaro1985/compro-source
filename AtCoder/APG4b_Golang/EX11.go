package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N, A int
	fmt.Scan(&N, &A)
	for i := 1; i <= N; i++ {
		var op string
		var B int
		fmt.Scan(&op, &B)
		ans := strconv.Itoa(i) + ":"
		switch op {
		case "+":
			A += B
			fmt.Println(ans + strconv.Itoa(A))
		case "-":
			A -= B
			fmt.Println(ans + strconv.Itoa(A))
		case "*":
			A *= B
			fmt.Println(ans + strconv.Itoa(A))
		case "/":
			if B == 0 {
				fmt.Println("error")
				return
			}
			A /= B
			fmt.Println(ans + strconv.Itoa(A))
		}
	}
}
