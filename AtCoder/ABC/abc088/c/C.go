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

func main() {
	c := make([][]int, 3)
	for i, _ := range c {
		c[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = nextInt()
		}
	}
	for i := 0; i <= 100; i++ {
		a := make([]int, 3)
		b := make([]int, 3)
		a[0] = i
		ok := true
		for j, _ := range c[0] {
			if c[0][j] >= i {
				b[j] = c[0][j] - i
			} else {
				ok = false
				break
			}
		}
		if ok != true {
			continue
		}
		a[1] = c[1][0] - b[0]
		for j, _ := range c[1] {
			if c[1][j]-b[j] != a[1] {
				ok = false
				break
			}
		}
		if ok != true {
			continue
		}
		a[2] = c[2][0] - b[0]
		for j, _ := range c[2] {
			if c[2][j]-b[j] != a[2] {
				ok = false
				break
			}
		}
		if ok != true {
			continue
		}
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
