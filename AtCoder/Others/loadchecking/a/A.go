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

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	fmt.Println(20)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
	fmt.Println(1)
}
