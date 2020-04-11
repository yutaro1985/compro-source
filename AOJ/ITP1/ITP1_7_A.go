package main

import "fmt"

func main() {
	for {
		var m, f, r int
		var score string
		fmt.Scan(&m, &f, &r)
		if m == -1 && f == -1 && r == -1 {
			break
		}
		if m == -1 || f == -1 || m+f < 30 {
			score = "F"
		} else if m+f >= 80 {
			score = "A"
		} else if m+f >= 65 {
			score = "B"
		} else if m+f >= 50 || r >= 50 {
			score = "C"
		} else if m+f >= 30 {
			score = "D"
		}
		fmt.Println(score)
	}
}
