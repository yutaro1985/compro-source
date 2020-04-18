package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	S := nextLine()
	var prev rune
	for _, Si := range S {
		if Si == prev {
			fmt.Println("Bad")
			return
		}
		prev = Si
	}
	fmt.Println("Good")
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
