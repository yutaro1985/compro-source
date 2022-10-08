// CumulativeSum はあるint型配列Aの累積和配列を返す
func CumulativeSum(A []int) []int {
	res := make([]int, 0, len(A)+1)
	res = append(res, 0)
	for _, v := range A {
		res = append(res, res[len(res)-1]+v)
	}
	return res
}