package main

import (
	"fmt"
)

func main() {
	for {
		var x string
		var sum int
		fmt.Scan(&x)
		if x == "0" {
			break
		}
		for _, num := range x {
			sum += int(num - '0')
		}
		fmt.Println(sum)
	}
}
