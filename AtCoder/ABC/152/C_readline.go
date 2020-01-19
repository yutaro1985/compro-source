package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
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

func main() {
	s1 := readLine()
	s2 := readLine()
	s2col := strings.Split(s2, " ")
	nums, _ := strconv.Atoi(s1)
	trueTotal := 0
	min := 1000000
	for i := 0; i < nums; i++ {
		num, _ := strconv.Atoi(s2col[i])
		if min >= num {
			trueTotal++
			min = num
		}
	}
	fmt.Println(trueTotal)
}
