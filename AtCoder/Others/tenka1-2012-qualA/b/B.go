package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1e6)

func main() {
	// readLineで行ごと読み込んでスペースを正規表現で,に置換
	// ただ、golangの正規表現は遅いと言われているので別解も考えてみる
	cn := readLine()
	reg := regexp.MustCompile(`\s+`)
	fmt.Println(reg.ReplaceAllString(cn, ","))
}

func readLine() string {
	buf := make([]byte, 0, 1e6)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
