package main

import "fmt"

func main() {
	var K, A, B int
	fmt.Scan(&K, &A, &B)
	// 1円を1回ビスケットを交換したときの初期値
	ans := B

	// 1回交換した後は、ビスケットA枚を1円に交換→1円をビスケットB枚に交換を繰り返す
	// ビスケット交換が間に合わない場合はビスケットを1枚ずつ増やす
	// 交換すると損する場合はひたすらビスケットを1枚ずつ増やす
	if A+2 < B && K > A {
		ans += ((K-(A+1))/2)*(B-A) + (K-(A+1))%2
		fmt.Println(ans)
	} else {
		fmt.Println(K + 1)
	}
}
