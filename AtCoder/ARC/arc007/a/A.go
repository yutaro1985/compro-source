package main

import (
	"fmt"
	"strings"
)

func main() {
	var X, s string
	fmt.Scan(&X, &s)
	fmt.Println(strings.Replace(s, X, "", -1))
}
