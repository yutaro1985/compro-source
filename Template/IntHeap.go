// IH golangの公式サンプルより
// https://xn--go-hh0g6u.com/pkg/container/heap/#example__intHeap
type IH []int

func (h IH) Len() int { return len(h) }

// 優先度付きキューとしてIntの値が大きい順にPOPしたいので大きい方をtrueに
func (h IH) Less(i, j int) bool { return h[i] > h[j] }
func (h IH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IH) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}