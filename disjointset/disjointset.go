package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// disjointSet is the implementation of the DisjointSet interface.
type disjointSet struct {
	parent map[int]int
	rank   map[int]int
}

// NewDisjointSet creates a new instance of a disjoint set.
func NewDisjointSet() DisjointSet {
	return &disjointSet{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

// FindSet finds the representative of the set containing x with path compression.
func (ds *disjointSet) FindSet(x int) int {
	if _, exists := ds.parent[x]; !exists {
		ds.parent[x] = x
		ds.rank[x] = 0
	}
	if ds.parent[x] != x {
		ds.parent[x] = ds.FindSet(ds.parent[x]) // Path compression
	}
	return ds.parent[x]
}

// UnionSet merges the sets containing x and y and returns the representative of the merged set.
func (ds *disjointSet) UnionSet(x, y int) int {
	xRoot := ds.FindSet(x)
	yRoot := ds.FindSet(y)

	if xRoot == yRoot {
		return xRoot
	}

	// Union by rank
	if ds.rank[xRoot] < ds.rank[yRoot] {
		ds.parent[xRoot] = yRoot
		return yRoot
	} else if ds.rank[xRoot] > ds.rank[yRoot] {
		ds.parent[yRoot] = xRoot
		return xRoot
	} else {
		ds.parent[yRoot] = xRoot
		ds.rank[xRoot]++
		return xRoot
	}
}
