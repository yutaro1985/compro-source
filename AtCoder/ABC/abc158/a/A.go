package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	if S == "AAA" || S == "BBB" {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
