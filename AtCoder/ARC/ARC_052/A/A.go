package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string
	fmt.Scan(&S)
	// 正規表現を使った抜き出し
	reg := regexp.MustCompile(`[0-9]+`)
	fmt.Println(reg.FindAllString(S, -1)[0])
}
