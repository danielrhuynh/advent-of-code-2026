package types

type DSU struct {
	Parent []int
	Size []int
}

func InitDSU(n int) DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return DSU {
		Parent: parent,
		Size: size,
	}
}

func (dsu *DSU) Find(x int) int {
	if dsu.Parent[x] != x {
		dsu.Parent[x] = dsu.Find(dsu.Parent[x])
	}
	return dsu.Parent[x]
}

func (dsu *DSU) Union(x, y int) (bool, int) {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)

	if rootX == rootY {
		return false, dsu.Size[rootX]
	}
	if dsu.Size[rootX] < dsu.Size[rootY] {
		rootX, rootY = rootY, rootX
	}
	dsu.Parent[rootY] = rootX
	dsu.Size[rootX] += dsu.Size[rootY]
	return true, dsu.Size[rootX]
}
