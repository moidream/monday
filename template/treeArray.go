package template

type NumArray struct {
	nums []int
	tree []int
}
func lowbit(x int) int{
	return x & -x
}
// 查询前缀和的方法
func (na *NumArray) query(x int) int {
	res := 0
	for i := x; i > 0; i -= lowbit(i) {
		res += na.tree[i]
	}
	return res
}
// 在树状数组 index 位置中增加值 val
func (na *NumArray) add(index int, val int) {
	for i := index; i < len(na.tree); i += lowbit(i) {
		na.tree[i] += val
	}
}
// 初始化「树状数组」，要默认数组是从 1 开始
func initTree(nums []int) NumArray {
	tree := make([]int, len(nums)+1)
	na := NumArray{nums, tree}
	for i := 0; i < len(nums); i ++ {
		na.add(i+1, nums[i])
	}
	return na
}


//使用树状数组
// 原有的值是 nums[i]，要使得修改为 val，需要增加 val - nums[i]
func (na *NumArray) Update(index int, val int)  {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}
//查找（left, right）之间的区间和
func (na *NumArray) SumRange(left int, right int) int {
	return na.query(right+1) - na.query(left)
}
