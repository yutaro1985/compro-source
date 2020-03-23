package main

import "fmt"

import "strconv"

func main() {
	var S int
	fmt.Scan(&S)
	h := S / 3600
	m := (S / 60) % 60
	s := S % 60
	fmt.Println(strconv.Itoa(h) + ":" + strconv.Itoa(m) + ":" + strconv.Itoa(s))
}
