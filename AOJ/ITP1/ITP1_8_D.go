package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, p string
	fmt.Scan(&s, &p)
	sd := s + s[:len(p)]
	if strings.Contains(sd, p) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
