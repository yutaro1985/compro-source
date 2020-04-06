package main

import "fmt"

func main() {
	// 他人の答えを参考にした
	var N, M, nextdisk int
	fmt.Scan(&N, &M)
	disk := make([]int, N+1)
	for i := range disk {
		disk[i] = i
	}
	for i := 0; i <= M; i++ {
		fmt.Scan(&nextdisk)
		for j, num := range disk {
			if num == nextdisk {
				disk[j] = disk[0]
				disk[0] = nextdisk
			}
		}
	}
	for i := 1; i <= N; i++ {
		fmt.Println(disk[i])
	}
}
