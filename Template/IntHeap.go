// IH golangの公式サンプルより
// https://cs.opensource.google/go/go/+/refs/tags/go1.22.1:src/container/heap/example_intheap_test.go
type IH []int

func InitIH() IH {
	ih := make(IH, 0)
	heap.Init(&ih)
	return ih
}

func (h IH) Len() int { return len(h) }

// 不等号の向きは都度調整する（デフォルトでは降順）
func (h IH) Less(i, j int) bool { return h[i] > h[j] }
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

func (h *IH) HPop() int {
	x := heap.Pop(h).(int)
	return x
}
