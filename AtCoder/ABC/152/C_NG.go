package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	var s1,s2 string
	s1 = nextLine()
	s2 = nextLine()
	nums, _ := strconv.Atoi(s1)
	slice := strings.Split(s2, " ")
	inputNum := stringToint(slice)

	truenums := 0
	// n番目の数値が最小値かどうかをチェックすることで十分
	min := inputNum[0]
	for i := 0; i < nums; i++ {
		if min >= inputNum[i] {
			min = inputNum[i]
			truenums++
		}
	}
	fmt.Println(truenums)
}
// 今回は未使用
// func Scanner() []string {
// 	sc := bufio.NewScanner(os.Stdin)
// 	var s string
// 	if sc.Scan() {
// 		s = sc.Text()
// 	}
// 	slice := strings.Split(s, " ")
// 	return slice
// }

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}
// 今回は未使用
// func max(a []int) int {
//     max := a[0]
//     for _, i := range a {
//         if i > max {
//             max = i
//         }
//     }
//     return max
// }

// func min(a []int) int {
//     min := a[0]
//     for _, i := range a {
//         if i < min {
//             min = i
//         }
//     }
//     return min
// }

