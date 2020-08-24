package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s1,s2 string
	if sc.Scan() {
        s1 = sc.Text()
    }
    if sc.Scan() {
        s2 = sc.Text()
    }
	slice1 := strings.Split(s1, " ")
	slice2 := strings.Split(s2, " ")
	inputNum1 := stringToint(slice1)
	inputNum2 := stringToint(slice2)
	sum := 0
	left := 0
	for i := 0; i<len(inputNum2); i++ {
		sum = sum + inputNum2[i]
	}
	if inputNum1[0]*inputNum1[2] - sum > inputNum1[1] {
		fmt.Println(-1)
	} else {
		left = inputNum1[2]*inputNum1[0] - sum
		if left <= 0 {
			fmt.Println(0)			
		} else {
			fmt.Println(left)
		}
	}
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
