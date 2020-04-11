package main

import (
	"fmt"
	"math"
)

func main() {
	//ただのテスト用
	fmt.Println(math.MaxInt64)
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(min(a, b))
}

// TODO #4 mathパッケージのMin、Maxを使った最小値、最大値がうまく行かなかった件について検証する
func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
