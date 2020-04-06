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
			// diskを一つずつ見ていって、次に出すディスクの番号と一致したらそこに現在のディスクを入れて、
			// 次のディスクはケースの外（配列の先頭）に出す
			if num == nextdisk {
				disk[j] = disk[0]
				disk[0] = nextdisk
				break
			}
		}
	}
	for i := 1; i <= N; i++ {
		fmt.Println(disk[i])
	}
}
