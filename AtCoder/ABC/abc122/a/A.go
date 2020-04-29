package main

import "fmt"

func main() {
	var b, ans string
	fmt.Scan(&b)
	switch b {
	case "A":
		ans = "T"
	case "C":
		ans = "G"
	case "G":
		ans = "C"
	case "T":
		ans = "A"
	}
	fmt.Println(ans)
}
