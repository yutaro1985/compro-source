// TopoligicalSort はトポロジカルソートを行う構造体
type TopologicalSort struct {
	Edges  [][]int
	order  []int
	degree []int
}

func newTopologicalSort(N int) *TopologicalSort {
	t := new(TopologicalSort)
	t.Edges = make([][]int, N)
	t.order = make([]int, 0)
	t.degree = make([]int, N)
	return t
}

func (t *TopologicalSort) addEdge(u, v int) {
	t.Edges[u] = append(t.Edges[u], v)
	t.degree[v]++
}

func (t *TopologicalSort) sort() bool {
	N := len(t.Edges)
	q := list.New()
	for i := 0; i < N; i++ {
		if t.degree[i] == 0 {
			q.PushBack(i)
		}
	}
	for q.Len() > 0 {
		v := q.Front().Value.(int)
		q.Remove(q.Front())
		for _, u := range t.Edges[v] {
			t.degree[u]--
			if t.degree[u] == 0 {
				q.PushBack(u)
			}
		}
		t.order = append(t.order, v)
	}
	return len(t.order) == N
}