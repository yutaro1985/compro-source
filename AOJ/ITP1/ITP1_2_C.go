package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 3)
	for i, _ := range nums {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	fmt.Println(nums[0], nums[1], nums[2])
}
