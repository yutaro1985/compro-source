package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
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
	H, W := nextInt(), nextInt()
	S := make([][]string, H)
	for i := range S {
		S[i] = strings.Split(nextLine(), "")
	}
	prevS := make([][]string, H)
	for i := range prevS {
		prevS[i] = make([]string, W)
	}
	recont := make([][]string, H)
	for i := range recont {
		recont[i] = make([]string, W)
	}
	dx := []int{1, 1, 0, -1, -1, -1, 0, 1, 0}
	dy := []int{0, -1, -1, -1, 0, 1, 1, 1, 0}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			isAllBlack := true
			if S[i][j] == "#" {
				for k := 0; k < len(dx); k++ {
					x := j + dx[k]
					y := i + dy[k]
					if x < 0 || x >= W || y < 0 || y >= H {
						continue
					}
					if S[y][x] != "#" {
						isAllBlack = false
						break
					}
				}

				if isAllBlack {
					for k := 0; k < len(dx); k++ {
						x := j + dx[k]
						y := i + dy[k]
						if x < 0 || x >= W || y < 0 || y >= H {
							continue
						}
						if prevS[y][x] != "#" {
							prevS[y][x] = "."
						}
					}
					prevS[i][j] = "#"
				} else {
					prevS[i][j] = "."
				}
			} else if prevS[i][j] != "#" {
				prevS[i][j] = "."
			}
			// fmt.Println(prevS)
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if prevS[i][j] == "#" {
				for k := 0; k < len(dx); k++ {
					x := j + dx[k]
					y := i + dy[k]
					if x < 0 || x >= W || y < 0 || y >= H {
						continue
					}
					recont[y][x] = "#"
				}
			} else if recont[i][j] != "#" {
				recont[i][j] = "."
			}
		}
	}
	// fmt.Println(S, prevS, recont)
	if reflect.DeepEqual(S, recont) {
		fmt.Println("possible")
		for _, s := range prevS {
			fmt.Println(strings.Join(s, ""))
		}
	} else {
		fmt.Println("impossible")
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
