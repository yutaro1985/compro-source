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
	// それぞれをbitで判定する
	three uint = 1
	five       = three << 1
	seven      = five << 1
	all        = seven | five | three
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 見様見真似実装
// https://drken1215.hatenablog.com/entry/2019/04/03/125400
// https://drken1215.hatenablog.com/entry/2020/05/04/190252

func main() {
	X := nextInt()
	var ans int
	dfs(0, 0, &ans, X)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func dfs(num int, used uint, ans *int, X int) {
	if num > X {
		return
	}
	if used == all {
		*ans++
	}
	// 実際に遷移を見るとわかりやすい
	// fmt.Println(num, used, *ans)
	// なんとなく3からのほうがしっくり来るのでそうしてるがどっちでもいい
	dfs(num*10+3, used|three, ans, X)
	dfs(num*10+5, used|five, ans, X)
	dfs(num*10+7, used|seven, ans, X)
}
