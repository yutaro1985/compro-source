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
	A := make([][]int, 3)
	for i := 0; i < 3; i++ {
		A[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			A[i][j] = nextInt()
		}
	}
	var N int
	N = nextInt()
	matrix := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for i := 0; i < N; i++ {
		b := nextInt()
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if A[j][k] == b {
					matrix[j][k] = 1
				}
				if matrix[j][0]+matrix[j][1]+matrix[j][2] == 3 {
					fmt.Println("Yes")
					return
				}
			}
		}
		if matrix[0][0]+matrix[1][0]+matrix[2][0] == 3 || matrix[0][1]+matrix[1][1]+matrix[1][1] == 3 || matrix[0][2]+matrix[1][2]+matrix[2][2] == 3 || matrix[0][0]+matrix[1][1]+matrix[2][2] == 3 || matrix[2][0]+matrix[1][1]+matrix[0][2] == 3 {
			fmt.Println("Yes")
			return
		}
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
