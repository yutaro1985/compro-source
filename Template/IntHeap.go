// IH golangの公式サンプルより
// https://xn--go-hh0g6u.com/pkg/container/heap/#example__intHeap
type IH []int

func (h IH) Len() int { return len(h) }

func (h IH) Less(i, j int) bool { return h[i] < h[j] }
func (h IH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IH) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IH) HPush(x interface{}) {
	heap.Push(h, x)
}

func (h *IH) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IH) HPop() interface{} {
	x := heap.Pop(h).(int)
	return x
}
