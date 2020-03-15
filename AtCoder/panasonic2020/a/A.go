package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51}
	var K int
	fmt.Scan(&K)
	fmt.Println(nums[K-1])
}
