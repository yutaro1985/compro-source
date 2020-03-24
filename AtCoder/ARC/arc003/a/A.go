package main

import "fmt"

import "strings"

func main() {
	var N float64
	var S string
	fmt.Scan(&N, &S)
	GPAs := strings.Split(S, "")
	scores := make(map[string]int)
	for _, GPA := range GPAs {
		scores[GPA]++
	}
	ans := float64(scores["A"]*4+scores["B"]*3+scores["C"]*2+scores["D"]) / N
	fmt.Println(ans)
}
