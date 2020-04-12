package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	N := nextInt()
	S := nextLine()
	var R, G, B, cnt int
	for i := 0; i < N; i++ {
		switch S[i] {
		case 'R':
			R++
		case 'G':
			G++
		case 'B':
			B++
		}
	}
	for i := 0; i < N-2; i++ {
		for j := 1; i+j+j < N; j++ {
			check := make([]byte, 3)
			check[0] = S[i]
			check[1] = S[i+j]
			check[2] = S[i+j+j]
			switch string(check) {
			case "RGB", "RBG", "GRB", "GBR", "BRG", "BGR":
				cnt++
			}
		}
	}
	fmt.Println(R*B*G - cnt)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
