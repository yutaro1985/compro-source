package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	var ans string
	var check uint
	var plus uint = 1
	minus := plus << 1
	both := plus | minus
	if A+B == C {
		check |= plus
	}
	if A-B == C {
		check |= minus
	}
	switch check {
	case plus:
		ans = "+"
	case minus:
		ans = "-"
	case both:
		ans = "?"
	default:
		ans = "!"
	}
	fmt.Println(ans)
}
