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
	H := nextInt()
	W := nextInt()
	S := make([]string, H)
	// 周囲1マスずつを定義
	dxy := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	for i := 0; i < H; i++ {
		S[i] = nextLine()
	}
	Srune := make([][]rune, H)
	for i := 0; i < H; i++ {
		Srune[i] = []rune(S[i])
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if Srune[i][j] == '#' {
				continue
			}
			var num int
			for d := 0; d < len(dxy); d++ {
				n := []int{j, i}
				n[0] = j + dxy[d][0]
				n[1] = i + dxy[d][1]
				if n[0] < 0 || W <= n[0] || n[1] < 0 || H <= n[1] {
					continue
				}
				if Srune[n[1]][n[0]] == '#' {
					num++
				}
			}
			Srune[i][j] = rune(num + '0')
		}
	}
	for i := 0; i < H; i++ {
		fmt.Println(string(Srune[i]))
	}
}

func searchMine([]int) {

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
