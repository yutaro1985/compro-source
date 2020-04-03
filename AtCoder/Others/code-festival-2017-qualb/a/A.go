package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	fmt.Println(S[:len(S)-8])
}
