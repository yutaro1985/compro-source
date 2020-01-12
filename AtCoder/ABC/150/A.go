package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	abyte := []byte(a)
	next := string(abyte[0]+1)
	fmt.Println(next)
}