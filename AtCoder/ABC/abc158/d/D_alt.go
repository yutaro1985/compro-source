package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://qiita.com/tnoda_/items/b503a72eac82862d30c6
// Scannerのサンプル
// https://baubaubau.hatenablog.com/entry/2017/11/17/214635
// byteで文字列操作を例

const (
	initialBufSize = 1e4
	maxBufsize     = 1e6 + 7
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	buf := make([]byte, initialBufSize)
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufsize)
}

func main() {
	sc.Scan()
	// Bytesでそのまま入れるとbufferが足りなくなる？
	S := append(make([]byte, 0, 1e6), sc.Bytes()...)

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())
	rev := false
	t := make([]byte, 0, 1e6)
	for i := 0; i < Q; i++ {
		sc.Scan()
		T, _ := strconv.Atoi(sc.Text())
		if T == 1 {
			rev = !rev
		} else {
			sc.Scan()
			F, _ := strconv.Atoi(sc.Text())
			F--
			sc.Scan()
			C := sc.Text()
			if rev {
				F = 1 - F
			}
			if F == 0 {
				t = append(t, C[0])
			} else {
				S = append(S, C[0])
			}
		}
	}
	t = ReverseByte(t)
	ans := append(t, S...)
	if rev {
		ans = ReverseByte(ans)
	}
	fmt.Println(string(ans))
}

func ReverseByte(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
