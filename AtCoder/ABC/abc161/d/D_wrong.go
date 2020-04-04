package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var K, ans int
	fmt.Scan(&K)
	count := 1
	var lunlun bool
	for i := 1; K >= count; i++ {
		fmt.Println(i, count)
		num := strconv.Itoa(i)
		if len(num) == 1 {
			ans = i
			count++
		} else {
			for j := 1; j < len(num); j++ {
				if int(num[j])-int(num[j-1]) > 1 {
					i += (10 - i%10) * int(math.Pow10(len(num)-2))
					lunlun = false
				}
				if abs(int(num[j]), int(num[j-1])) < 1 {
					lunlun = false
					break
				} else {
					lunlun = true
				}
			}
			if lunlun {
				ans = i
				count++
			}
		}
		// if len(num) == 1 || islunlun(num) {
		// 	ans = i
		// 	count++
		// } else if num[i]-num[i-1] > 1 {
		// 	// i += (10 - i%10) * int(math.Pow10(len(num)-1))
		// 	i += (10 - i%10)
		// }
	}
	fmt.Println(ans)
}

func islunlun(num string) bool {
	if len(num) != 1 {
		for i := 1; i < len(num); i++ {
			if abs(int(num[i]), int(num[i-1])) > 1 {
				return false
			}
		}
	} else {
		return true
	}
	return true
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
