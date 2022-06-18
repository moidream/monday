package template

import (
	"container/heap"
	"sort"
)

//sort.IntSlice
type IS struct{
	sort.IntSlice
}
func (is *IS)Push(x interface{}){
	is.IntSlice = append(is.IntSlice, x.(int))
}
func (is *IS)Pop()interface{}{
	e := is.IntSlice[is.Len()-1]
	is.IntSlice = is.IntSlice[:is.Len()-1]
	return e
}
func oprate() {
	h:=&IS{}
	heap.Init(h)
	heap.Push(h, 1)
	_ = heap.Pop(h).(int)
}


//其他写法
type IHeap [][2]int
func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
