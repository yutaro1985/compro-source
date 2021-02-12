q := list.New()
// 初期位置をPushBack
q.PushBack(Position{})
for q.Len() > 0 {
	P := q.Front().Value.(Position)
	q.Remove(q.Front())
	for _, v := range d {
		nextP := Position{P.H + v.H, P.W + v.W}
		if nextP.H < 0 || nextP.H >= H || nextP.W < 0 || nextP.W >= W {
			continue
		}
		if seen[nextP.H][nextP.W] {
			continue
		}
		if dist[nextP.H][nextP.W] != -1 {
			dist[nextP.H][nextP.W] = dist[P.H][P.W] + 1
			seen[nextP.H][nextP.W] = true
			q.PushBack(nextP)
		}
	}
}