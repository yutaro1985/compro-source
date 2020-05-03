package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	check := map[rune]bool{}
	for _, s := range S {
		if _, exist := check[s]; exist {
			fmt.Println("no")
			return
		} else {
			check[s] = true
		}
	}
	fmt.Println("yes")
}
