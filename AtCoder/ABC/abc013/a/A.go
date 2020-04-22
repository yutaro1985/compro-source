package main

import "fmt"

func main() {
	var X string
	var ans int
	fmt.Scan(&X)
	switch X {
	case "A":
		ans = 1
	case "B":
		ans = 2
	case "C":
		ans = 3
	case "D":
		ans = 4
	case "E":
		ans = 5
	}
	fmt.Println(ans)
}
