// SegmentTree はSegment木の実装
// こちらのもの: https://zenn.dev/johniel/articles/f1028e37f91489
type SegmentTree struct {
	data []interface{}
	n    int
	e    interface{}
	op   func(interface{}, interface{}) interface{}
}

func NewSegmentTree(n int, e interface{}, op func(interface{}, interface{}) interface{}) *SegmentTree {
	segtree := new(SegmentTree)
	segtree.e = e
	segtree.op = op
	segtree.n = 1
	for segtree.n < n {
		segtree.n *= 2
	}
	segtree.data = make([]interface{}, segtree.n*2-1)
	for i := 0; i < segtree.n*2-1; i++ {
		segtree.data[i] = segtree.e
	}
	return segtree
}

func (segtree *SegmentTree) Update(idx int, x interface{}) {
	idx += segtree.n - 1
	segtree.data[idx] = x
	for 0 < idx {
		idx = (idx - 1) / 2
		segtree.data[idx] = segtree.op(segtree.data[idx*2+1], segtree.data[idx*2+2])
	}
}

func (segtree *SegmentTree) query(begin, end, idx, a, b int) interface{} {
	if b <= begin || end <= a {
		return segtree.e
	}
	if begin <= a && b <= end {
		return segtree.data[idx]
	}
	v1 := segtree.query(begin, end, idx*2+1, a, (a+b)/2)
	v2 := segtree.query(begin, end, idx*2+2, (a+b)/2, b)
	return segtree.op(v1, v2)
}

func (segtree *SegmentTree) Query(begin, end int) interface{} {
	return segtree.query(begin, end, 0, 0, segtree.n)
}

// NewSegmentTree(N, 0, func(a, b interface{}) interface{} { return a.(int) ^ b.(int) })