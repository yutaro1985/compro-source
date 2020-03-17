package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	command := nextLine()
	//TODO #5 このケースでも通ってしまうが、本来はダメらしいので正しい解き方のパターンを作る
	L := []string{"AA", "AB", "AX", "AY", "BA", "BB", "BX", "BY", "XA", "XB", "XX", "XY", "YA", "YB", "YX", "YY"}
	R := []string{"AA", "AB", "AX", "AY", "BA", "BB", "BX", "BY", "XA", "XB", "XX", "XY", "YA", "YB", "YX", "YY"}
	min := N + 1
	for i, L := range L {
		for j, R := range R {
			if i < j {
				shortcutcommand := strings.Replace(command, L, "C", -1)
				shortcutcommand = strings.Replace(shortcutcommand, R, "D", -1)
				if min > len(shortcutcommand) {
					min = len(shortcutcommand)
				}
			}
		}
	}
	fmt.Println(min)
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
