// BIT = fenwick tree はBinary Indexed Treeの実装
type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	b := new(BIT)
	b.v = make([]int, n)
	return b
}

func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

func (b BIT) sum(a int) int {
	res := 0
	for i := a + 1; i > 0; i -= i & -i {
		res += b.v[i-1]
	}
	return res
}

func (b BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		return b.sum(y) - b.sum(x-1)
	}
}

func (b BIT) lowerbound(x int) int {
	idx, k := 0, 1
	for k < len(b.v) {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k-1 < len(b.v) && b.v[idx+k-1] < x {
			x -= b.v[idx+k-1]
			idx += k
		}
	}
	return idx
}
