package shared

type Point3D struct {
	X, Y, Z int
}

func EuclideanDistance3D(p1, p2 Point3D) int {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	// skip math.sqrt as it is sufficient to return the squared value most of the time
	return dx*dx + dy*dy + dz*dz
}

func Map[T, V any](ts []T, fn func(T) (V, error)) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i], _ = fn(t)
	}
	return result
}

type UnionFind struct {
	parent []int
	Size   []int
}

func NewUnionFind(size int) *UnionFind {
	return &UnionFind{
		parent: make([]int, size),
		Size:   make([]int, size),
	}
}

func (uf *UnionFind) Add(x int) {
	uf.parent[x] = x
	uf.Size[x] = 1
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}

	// smaller merge into the bigger
	if uf.Size[rootX] > uf.Size[rootY] {
		uf.parent[rootY] = rootX
		uf.Size[rootX] += uf.Size[rootY]
		uf.Size[rootY] = 0
	} else {
		uf.parent[rootX] = rootY
		uf.Size[rootY] += uf.Size[rootX]
		uf.Size[rootX] = 0
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}
