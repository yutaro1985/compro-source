// ArithSum は初項a、公差d、項数nの等差数列の和を返す
func ArithSum(a, d, n int) int {
	var res int
	res = n * (2*a + (n-1)*d) / 2
	return res
}
