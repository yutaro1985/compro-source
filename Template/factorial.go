// マイナスの場合は考慮していない
func factorialmod(a, m int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return (factorialmod(a-1, m) * a) % m
	}
}