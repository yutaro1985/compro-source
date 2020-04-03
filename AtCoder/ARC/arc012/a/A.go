package main

import "fmt"

func main() {
	var day string
	var rest int
	fmt.Scan(&day)
	switch day {
	case "Monday":
		rest = 5
	case "Tuesday":
		rest = 4
	case "Wednesday":
		rest = 3
	case "Thursday":
		rest = 2
	case "Friday":
		rest = 1
	}
	fmt.Println(rest)
}
