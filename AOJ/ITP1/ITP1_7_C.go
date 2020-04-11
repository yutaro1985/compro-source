package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	sumcol := make([]int, c+1)
	for i := 0; i < r; i++ {
		var sum int
		for j := 0; j < c; j++ {
			var n int
			fmt.Scan(&n)
			sum += n
			sumcol[j] += n
			fmt.Printf("%d ", n)
			if j == c-1 {
				sumcol[j+1] += sum
				fmt.Println(sum)
			}
		}
		if i == r-1 {
			for k, sums := range sumcol {
				if k != len(sumcol)-1 {
					fmt.Printf("%d ", sums)
				} else {
					fmt.Println(sums)
				}
			}
		}
	}
}
