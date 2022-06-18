package template

type UnionFind struct {
	parent []int
	cnt int
}
func Init(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i ++ {
		parent[i] = i
	}
	return &UnionFind{
		parent : parent,
		cnt : n,
	}
}
func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[pi] = pj
		u.cnt --
	}
}
func (u *UnionFind) Find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] != i { //路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}
