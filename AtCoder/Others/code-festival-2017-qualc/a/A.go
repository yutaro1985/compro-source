package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	if strings.Contains(S, "AC") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
