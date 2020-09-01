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
	N := nextInt()
	M := nextInt()
	s := make([]int, M)
	c := make([]int, M)
	ans := -1
	for i := 0; i < M; i++ {
		s[i] = nextInt() - 1
		c[i] = nextInt()
	}
	// 計算量が多くないので総当りでチェックできる
	for i := 0; i < 1000; i++ {
		nums := strconv.Itoa(i)
		check := true
		// 桁数をまずはチェック。数値をstringに変換してlenで桁数が出せる
		if len(nums) == N {
			//  条件を一つずつ確認する。M=0の場合は何も行われないでそのままcheck == true
			for j := 0; j < M; j++ {
				// 確認中の数値がそれぞれの条件に一致するかをチェックする。
				num, _ := strconv.Atoi(string(nums[s[j]]))
				if num != c[j] {
					check = false
				}
			}
		} else {
			check = false
		}
		if check == true {
			ans = i
			fmt.Println(ans)
			return
		}
	}
	fmt.Println(ans)
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
