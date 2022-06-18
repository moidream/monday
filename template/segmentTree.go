package template

//每个节点表示数组区间[l, r]间的和
type seg []struct{ l, r, sum int }

//o为线段树当前节点，节点下标从0开始
//build(0, 0, len(nums)-1, nums)表示从index为0的节点开始建立数组的线段树
func (t seg) build(o, l, r int, nums []int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = nums[l]
		return
	}
	m := (l + r) / 2
	t.build(o*2+1, l, m, nums)
	t.build(o*2+2, m+1, r, nums)
	t[o].sum = t[o*2+1].sum + t[o*2+2].sum
}

// 将数组 idx 上的元素值增加 val
func (t seg) add(o, idx, val int) {
	if t[o].l == t[o].r {
		t[o].sum += val
		return
	}
	m := (t[o].l + t[o].r) / 2
	if idx <= m {
		t.add(o*2+1, idx, val)
	} else {
		t.add(o*2+2, idx, val)
	}
	lo, ro := t[o*2+1], t[o*2+2]
	t[o].sum = lo.sum + ro.sum
}

// 返回区间数组 [l,r] 内的元素和
func (t seg) querySum(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	m := (t[o].l + t[o].r) / 2
	if l <= m {
		res += t.querySum(o*2+1, l, r)
	}
	if r > m {
		res += t.querySum(o*2+2, l, r)
	}
	return
}
