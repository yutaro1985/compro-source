// TopoligicalSort はトポロジカルソートを行う構造体
type TopologicalSort struct {
	Edges [][]int
	order []int
}

func newTopologicalSort(N int) *TopologicalSort {
	t := new(TopologicalSort)
	t.Edges = make([][]int, N)
	t.order = make([]int, 0)
	return t
}

func (t *TopologicalSort) addEdge(u, v int) {
	t.Edges[u] = append(t.Edges[u], v)
}

func (t *TopologicalSort) sort() bool {
	N := len(t.Edges)
	seen := make([]int, N)
	order := make([]int, 0)
	var dfs func(G [][]int, v int) bool
	dfs = func(G [][]int, v int) bool {
		seen[v] = 1
		for _, nv := range G[v] {
			switch seen[nv] {
			case 0:
				if !dfs(G, nv) {
					return false
				}
			case 1:
				return false
			case 2:
				continue
			}
		}
		order = append(order, v)
		seen[v] = 2
		return true
	}
	for v := 0; v < N; v++ {
		if seen[v] == 0 && !dfs(t.Edges, v) {
			return false
		}
	}
	for i := 0; i < len(order); i++ {
		t.order = append(t.order, order[N-1-i])
	}
	return true
}