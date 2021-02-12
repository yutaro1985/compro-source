q := list.New()
q.PushBack(p)
dist[p] = 0
for q.Len() > 0 {
	cur := q.Front().Value.(int)
	q.Remove(q.Front())
	for _, next := range G[cur] {
		if dist[next] != INF {
			continue
		}
		dist[next] = dist[cur] + 1
		q.PushBack(next)
	}
}
