package main

import "fmt"

func main() {
	var inputNum int
	fmt.Scan(&inputNum)
	matches := inputNum * (inputNum - 1) / 2
	fmt.Println(matches)
}
