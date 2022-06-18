package template

import (
	"container/heap"
	"math"
)

//邻接表存储有向图
type edge struct {to, wt int}
func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt32
	}
	dis[start] = 0
	h := hp{pair{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		v := p.v
		if p.dis > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newDis := dis[v] + e.wt; newDis < dis[w] {
				dis[w] = newDis
				heap.Push(&h, pair{w, newDis})
			}
		}
	}
	return dis
}

type pair struct {v, dis int}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].dis < h[j].dis}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(v interface{}) {*h = append(*h, v.(pair))}
func (h *hp) Pop() (v interface{}) {
	a := *h
	*h, v = a[:len(a)-1], a[len(a)-1]
	return
}
