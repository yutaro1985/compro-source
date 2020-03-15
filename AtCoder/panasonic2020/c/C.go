package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	// 誤差により小数同士の比較は想定通りにならないケースがあるため、条件を整数で比較する形にする
	// c-a-bがマイナスの場合は不等号の向きが変わってしまうので場合分けが必要だが、今回のケースだと成立しないケースしかないのでNo
	condition := (c-a-b)*(c-a-b) - 4*a*b
	if condition > 0 && c > a+b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
