package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000009
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 嫌いな数字リストを好きな数字リストにしてうまくできないか考えたが思い浮かばなかった
func main() {
	N := nextInt()
	K := nextInt()
	D := make([]int, K)
	// d := nextInt()
	// count := 1
	// for i := 0; i < 10; i++ {
	// 	if d != i {
	// 		L = append(L, i)
	// 	} else if i != 9 && count != K {
	// 		d = nextInt()
	// 		count++
	// 	}
	// }
	for i := range D {
		D[i] = nextInt()
	}
	for i := 0; i < N*10; i++ {
		ok := true
		for _, num := range strconv.Itoa(i) {
			for _, d := range D {
				if num == rune(d+'0') {
					ok = false
				}
			}
		}
		if ok && i >= N {
			fmt.Println(i)
			return
		}
	}
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
