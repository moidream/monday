package template

import "math/rand"

/*二分查找, 求 >= , 返回 left
func test() {
	a := []int{1, 2, 4, 6, 8, 10}
	target := 11
	res1 := binarySearch(a, target)
	res2 := sort.SearchInts(a, target)
	res3 := sort.Search(len(a), func(i int) bool {
		return a[i] >= target
	})
	fmt.Println(res1, res2, res3)
}
*/
func binarySearch(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l+r) / 2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

//单调栈，单调递增栈会剔除波峰，留下波谷
//当前项向左找第一个比自己大的位置 —— 从左向右维护一个单调递减栈
//当前项向右找第一个比自己小的位置 —— 从右向左维护一个单调递增栈
func monotonicStack(nums []int) ([]int, []int) {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	//单调栈
	for i, _ := range right {
		right[i] = n
	}
	stack := []int{}
	//一次求左右第一个比当前小的数，right数组并不准确，边界有误
	for i := 0; i < n; i ++ {
		//当前项向左找第一个比自己小的数，从左到右维护单调递增栈
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return left, right
}

//快速排序
func quickSort(nums []int, left, right int) []int {
	if left >= right {
		return nums
	}
	p := rand.Intn(right-left+1)+left
	nums[left], nums[p] = nums[p], nums[left]
	pivot := nums[left]
	j := left
	for i := left+1; i <= right; i ++ {
		if nums[i] < pivot {
			j ++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	//结束循环后，j所指向即为pivot应在位置
	nums[left], nums[j] = nums[j], nums[left]
	quickSort(nums, left, j-1)
	quickSort(nums, j+1, right)
	return nums
}

//手建大根堆， 求数组中第 K 大的数
func findKthLargest(nums []int, k int) int {
	heapsize := len(nums)
	buildHeap(nums, heapsize)
	for i := len(nums)-1; i >= len(nums)-k+1; i -- {
		nums[0], nums[i] = nums[i], nums[0]
		heapsize--
		maxifyHeap(nums, 0, heapsize)
	}
	return nums[0]
}
func buildHeap(heap []int, heapsize int) {
	for i := heapsize/2; i >= 0; i -- {
		maxifyHeap(heap, i, heapsize)
	}
}
func maxifyHeap(heap []int, i int, heapsize int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < heapsize && heap[left] > heap[largest] {
		largest = left
	}
	if right < heapsize && heap[right] > heap[largest] {
		largest = right
	}
	if largest != i {
		heap[largest], heap[i] = heap[i], heap[largest]
		maxifyHeap(heap, largest, heapsize)
	}
}

// K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	//记录长度
	dummyHead := &ListNode{Next:head}
	prev := dummyHead
	for head != nil {
		tail := prev
		for i := 0; i < k; i ++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}
		next := tail.Next
		newHead, newTail := reverse(head, tail)
		prev.Next = newHead
		newTail.Next = next
		prev = newTail
		head = next
	}
	return dummyHead.Next
}
func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

// dfs 回溯，全排列
func permute(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	vis := map[int]bool{}
	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == n {
			res = append(res, append([]int{}, path...))
			return
		}
		for _, num := range nums {
			if !vis[num] {
				vis[num] = true
				path = append(path, num)
				dfs(path)
				path = path[:len(path)-1]
				vis[num] = false
			}
		}
	}
	dfs([]int{})
	return res
}

//双指针， 滑动窗口
func countSubarrays(nums []int, k int) int {
	res, left := 0, 0
	sum := 0
	for right, num := range nums {
		sum += num
		for sum * (right-left+1) >= k {
			sum -= nums[left]
			left ++
		}
		res += right - left + 1
	}
	return res
}


/*
首先是背包分类的模板
	1、0/1背包：外循环nums,内循环target,target倒序且target>=nums[i];
	2、完全背包：外循环nums,内循环target,target正序且target>=nums[i];
	3、组合背包：外循环target,内循环nums,target正序且target>=nums[i];
	4、分组背包：这个比较特殊，需要三重循环：外循环背包bags,内部两层循环根据题目的要求转化为1,2,3三种背包类型的模板。
然后是问题分类的模板：
	1、最值问题: dp[i] = max/min(dp[i], dp[i-nums]+1)或dp[i] = max/min(dp[i], dp[i-num]+nums);
	2、存在问题(bool)：dp[i]=dp[i]||dp[i-num];
	3、组合问题：dp[i]+=dp[i-num];
*/
//零钱兑换，完全背包最值问题
func coinChange(coins []int, amount int) int {
	//dp【i】表示组成金额i所需最小硬币数
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount+1
	}
	dp[0] = 0
	for i := 1; i <= amount; i ++ {
		//枚举硬币j
		for j := 0; j < len(coins); j ++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]] + 1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

//随机数生成
func rand10() int {
	x := 50
	for x > 40 {
		x = 7 * (rand7()-1) + rand7() //1-49
	}
	return x%10 + 1
}
func rand7() int {
	return 1+rand.Intn(7)
}

//搜寻旋转排序数组||，有重复值
func search(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid-1
			} else {
				left = mid+1
			}
		} else if nums[left] > nums[mid] {
			if nums[mid] < target && target <= nums[right] {
				left = mid+1
			} else {
				right = mid -1
			}
		} else {
			left++
		}
	}
	return false
}


//预测赢家，石子游戏，一个数组取头尾，得分为所取值或余下值，若数组长度为偶数，先手必胜
func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = nums[i]
	}
	for i := n-2; i >= 0; i -- {
		for j := i+1; j < n; j ++ {
			dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[n-1] >= 0
}

//区间交集，判断最小末端点，每个末端点只可能与一个区间相交
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		low := max(firstList[i][0], secondList[j][0])
		high := min(firstList[i][1], secondList[j][1])
		if low <= high {
			res = append(res, []int{low, high})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

//差分数组，大小为原数组+1，前i项和即为原数组第i项
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	res := make([]int, n)
	for _, booking := range bookings {
		first, last, seats := booking[0]-1, booking[1]-1, booking[2]
		diff[first] += seats
		diff[last+1] -= seats
	}
	sum := 0
	for i := 0; i < n; i ++ {
		sum += diff[i]
		res[i] = sum
	}
	return res
}

// 0-1BFS，双端队列
func minimumObstacles(grid [][]int) int {
	type pair1 struct {
		x, y int
	}
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = m*n
		}
	}
	dir := []int{1, 0, -1, 0, 1}
	dis[0][0] = 0
	//双端队列，两个切片
	q := [2][]pair1{{{0, 0}}}
	for len(q[0]) > 0 || len(q[1]) > 0 {
		p := pair1{}
		if len(q[0]) > 0 {
			p = q[0][len(q[0])-1]
			q[0] = q[0][:len(q[0])-1]
		} else {
			p = q[1][0]
			q[1] = q[1][1:]
		}
		for i := 0; i < 4; i ++ {
			x, y := p.x + dir[i], p.y + dir[i+1]
			if x >= 0 && x < m && y >= 0 && y < n {
				g := grid[x][y]
				if dis[p.x][p.y] + g < dis[x][y] {
					dis[x][y] = dis[p.x][p.y] + g
					q[g] = append(q[g], pair1{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1]
}


