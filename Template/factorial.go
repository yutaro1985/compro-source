// マイナスの場合は考慮していない
func factorial(a int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return factorial(a-1) * a
	}
}