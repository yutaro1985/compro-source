package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	contname := S[0 : len(S)-4]
	y := S[len(S)-4 : len(S)]
	yi, _ := strconv.Atoi(y)
	y = strconv.Itoa(yi + 1)
	fmt.Println(contname + y)
}
