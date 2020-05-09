package main

import (
	"fmt"
)

func main() {
	var X int
	fmt.Scan(&X)
	for i := 0; ; i++ {
		if i*i*i*i == X {
			fmt.Println(i)
			return
		}
	}
}
