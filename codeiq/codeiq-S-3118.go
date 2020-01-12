package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var num [][]int
	for sc.Scan() {
		num = append(num, (stringToint(strings.Split(sc.Text(), ""))))
	}

}
func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}

func routeCalc(num [][]int) int {

}
