// CoordinateShrink は座標圧縮の対応させるスライスを返す関数
func CoordinateShrink(A []int) []int {
	N := len(A)
	B := make([]int, N)
	copy(B, A)
	sort.Ints(B)
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = sort.SearchInts(B, A[i])
	}
	return res
}
