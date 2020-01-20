package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, _ := strconv.Atoi(s1)
	s2 := readLine()
	an := strings.Split(s2, " ")
	Alice := 0
	Bob := 0
	sort.Strings(an)
	// 逆順ソート
	// sort.Sort(sort.Reverse(sort.StringSlice(an)))
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(an[n-i-1])
		switch i % 2 {
		case 0:
			Alice += num
		case 1:
			Bob += num
		default:
			fmt.Println("error")
			break
		}
		fmt.Println(i, num)
	}
	fmt.Println(Alice - Bob)
}
