// Dijkstra はダイクストラ法で各頂点への距離を求める
func Dijkstra(s int, Edges [][]Edge, dist []int) {
	pq := InitPQ()
	push := func(v, d int, dist []int) {
		if dist[v] <= d {
			return
		}
		dist[v] = d
		pq.HPush(Item{v, d, 0})
	}
	push(s, 0, dist)
	for pq.Len() > 0 {
		r := pq.HPop()
		if dist[r.to] != r.priority {
			continue
		}
		for _, v := range Edges[r.to] {
			push(v.to, r.priority+v.C, dist)
		}
	}
}

type Edge struct {
	to, C int
}

type Item struct {
	to, priority, index int
}

type PQ []*Item

func InitPQ() PQ {
	pq := make(PQ, 0)
	heap.Init(&pq)
	return pq
}

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) HPush(x Item) {
	heap.Push(pq, &x)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PQ) HPop() *Item {
	item := heap.Pop(pq).(*Item)
	return item
}
