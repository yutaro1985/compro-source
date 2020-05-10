package main

import "fmt"

func main() {
	var a, b, c bool
	a = true
	b = false
	c = true
	if a {
		fmt.Printf("%s", "At")
	} else {
		fmt.Printf("%s", "Yo")
	}

	if !a && b {
		fmt.Printf("%s", "Bo")
	} else if !b || c {
		fmt.Printf("%s", "Co")
	}

	if a && b && c {
		fmt.Printf("%s", "foo")
	} else if true && false {
		fmt.Printf("%s", "yeah!")
	} else if !a || c {
		fmt.Printf("%s", "der")
	}
	fmt.Println()
}
