// LIS はintのスライスを与えたとき最長増加部分列(Longest Increasing Subsequence)の長さを返す関数
func LIS(a []int) int {
	N := len(a)
	lis := make([]int, N)
	for i := 0; i < N; i++ {
		lis[i] = INF
	}
	for i := 0; i < N; i++ {
		idx := sort.SearchInts(lis, a[i])
		lis[idx] = a[i]
	}
	res := sort.SearchInts(lis, INF)
	return res
}
