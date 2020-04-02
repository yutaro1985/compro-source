package main

import "fmt"

func main() {
	// 6つの数字の中に、1回しか登場しない数字が2つ、2回登場する数字が2つになるような組み合わせだとOKということに着目
	road := make([]int, 5)
	check := make([]int, 6)
	for i := 0; i < 3; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		road[a]++
		road[b]++
	}
	for i := 1; i <= 4; i++ {
		check[road[i]-1]++
	}
	if check[0] == 2 && check[1] == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
