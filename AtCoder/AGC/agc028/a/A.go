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
	N, M := nextInt(), nextInt()
	S, T := nextLine(), nextLine()
	Smap, Tmap := map[int]byte{}, map[int]byte{}
	LCM := lcmof2numbers(N, M)
	for i := 0; i < N; i++ {
		Smap[LCM/N*i+1] = S[i]
	}
	for i := 0; i < M; i++ {
		Tmap[LCM/M*i+1] = T[i]
	}
	for i := 0; LCM/M*i+1 <= LCM; i++ {
		k := LCM/M*i + 1
		Sk, ifExist := Smap[k]
		if ifExist && Tmap[k] != Sk {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(LCM)
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

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}

func lcmof2numbers(a int, b int) int {
	return a / gcdof2numbers(a, b) * b
}
