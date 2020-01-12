package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	s := Scanner(sc)
	inputNum := stringToint(s)
	row := make([]int, inputNum[0])
	sumOfSales := 0
	totalSales := 0
	for i := 0; i < inputNum[1]; i++ {
		s = Scanner(sc)
		row = stringToint(s)
		for j := range row {
			sumOfSales += row[j]
		}
		if sumOfSales > 0 {
			totalSales += sumOfSales
		}
		sumOfSales = 0
	}
	fmt.Println(totalSales)
}

func Scanner(sc *bufio.Scanner) []string {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	slice := strings.Split(s, " ")
	return slice
}

func stringToint(s []string) []int {
	f := make([]int, len(s))
	for n := range s {
		f[n], _ = strconv.Atoi(s[n])
	}
	return f
}
