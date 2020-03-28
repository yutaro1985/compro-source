package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	building := make([][][]int, 4)
	for i, _ := range building {
		building[i] = make([][]int, 3)
		for j, _ := range building[i] {
			building[i][j] = make([]int, 10)
		}
	}
	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scan(&b, &f, &r, &v)
		building[b-1][f-1][r-1] += v
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			residents := ""
			for k := 0; k < 10; k++ {
				residents += " " + strconv.Itoa(building[i][j][k])
			}
			fmt.Println(residents)
		}
		if i != 3 {
			fmt.Println("####################")
		}
	}
}
