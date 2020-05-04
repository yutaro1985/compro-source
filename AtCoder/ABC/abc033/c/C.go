package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	formula := strings.Split(S, "+")
	ans := len(formula)
	for _, f := range formula {
		if f == "0" || strings.Contains(f, "0*") || strings.Contains(f, "*0") {
			ans--
		}
	}
	fmt.Println(ans)
}
