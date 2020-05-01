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
	colors := map[int]int{}
	for i := 0; i < N; i++ {
		colors[checkColor(nextInt())]++
	}
	var min, max int
	if _, exist := colors[9]; exist {
		if len(colors) > 1 {
			max = len(colors) + colors[9] - 1
			min = len(colors) - 1
		} else {
			max = colors[9]
			min = 1
		}
	} else {
		max = len(colors)
		min = max
	}
	fmt.Println(min, max)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func checkColor(a int) int {
	colornum := 0
	switch {
	case a < 400:
		colornum = 1
	case 400 <= a && a < 800:
		colornum = 2
	case 800 <= a && a < 1200:
		colornum = 3
	case 1200 <= a && a < 1600:
		colornum = 4
	case 1600 <= a && a < 2000:
		colornum = 5
	case 2000 <= a && a < 2400:
		colornum = 6
	case 2400 <= a && a < 2800:
		colornum = 7
	case 2800 <= a && a < 3200:
		colornum = 8
	case 3200 <= a:
		colornum = 9
	}
	return colornum
}
