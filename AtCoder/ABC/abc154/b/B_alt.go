package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	// strings.Repeatを使ったパターン
	fmt.Println(strings.Repeat("x", len(S)))
}
