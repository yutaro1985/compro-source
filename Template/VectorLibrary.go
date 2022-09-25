type Vertex struct {
	X, Y int
}

func newVector(from, to Vertex) Vector {
	return Vector{to.X - from.X, to.Y - from.Y}
}

type Vector struct {
	X, Y int
}

func (a *Vector) add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}

func (a *Vector) sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

func (a *Vector) addSelf(b Vector) {
	*a = a.add(b)
}

func (a *Vector) subSelf(b Vector) {
	*a = a.sub(b)
}

func (a *Vector) inner(b Vector) int {
	return a.X*b.X + a.Y*b.Y
}

func (a *Vector) cross(b Vector) int {
	return a.X*b.Y - a.Y*b.X
}