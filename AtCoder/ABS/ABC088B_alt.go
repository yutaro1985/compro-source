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
	an := stringToint(strings.Split(s2, " "))
	Total := 0
	sort.Ints(an)
	// 逆順ソート
	// sort.Sort(sort.Reverse(sort.StringSlice(an)))

	for i := 0; i < n; i++ {
		num := an[n-i-1]
		switch i % 2 {
		case 0:
			Total += num
		case 1:
			Total -= num
		default:
			fmt.Println("error")
			break
		}
	}
	fmt.Println(Total)
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}

// こちらのほうがスマート
//　https://qiita.com/Suzuki_Cecil/items/3620ad0445dc32d3d02a#%E7%AC%AC-%EF%BC%96-%E5%95%8F-abc-088-b---card-game-for-two-200-%E7%82%B9
