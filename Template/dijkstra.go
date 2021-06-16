// Dijkstra はダイクストラ法で各頂点への距離を求める
func Dijkstra(s int, Edges [][]Edge, dist []int) {
	pq := make(PQ, 0)
	heap.Init(&pq)
	push := func(v, d int, dist []int) {
		if dist[v] <= d {
			return
		}
		dist[v] = d
		heap.Push(&pq, &Item{v, d, 0})
	}
	push(s, 0, dist)
	for pq.Len() > 0 {
		r := heap.Pop(&pq).(*Item)
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

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}