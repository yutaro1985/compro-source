package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	keyboards := "WBWBWWBWBWBWWBWBWWBWBWBW"
	keycheck := map[int]string{
		0:  "Do",
		1:  "Do",
		2:  "Re",
		3:  "Re",
		4:  "Mi",
		5:  "Fa",
		6:  "Fa",
		7:  "So",
		8:  "So",
		9:  "La",
		10: "La",
		11: "Si",
	}
	fmt.Println(keycheck[strings.Index(keyboards, S)])
}
