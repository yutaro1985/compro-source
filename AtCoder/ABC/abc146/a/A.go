package main

import "fmt"

func main() {
	var S string
	var ans int
	fmt.Scan(&S)
	switch S {
	case "SUN":
		ans = 7
	case "MON":
		ans = 6
	case "TUE":
		ans = 5
	case "WED":
		ans = 4
	case "THU":
		ans = 3
	case "FRI":
		ans = 2
	case "SAT":
		ans = 1
	}
	fmt.Println(ans)
}
