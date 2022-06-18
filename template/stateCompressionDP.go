package template

func distributeCookies(cookies []int, k int) int {
	n := len(cookies)
	//sum表示某种分配状态下的结果值
	sum := make([]int, 1 << n)
	for i := 1; i < 1<<n; i ++ {
		for j := 0; j < n; j ++ {
			// i第j位为1，表示j分配
			if (i>>j & 1) == 1{
				sum[i] += cookies[j]
			}
		}
	}
	//dp[i][j]表示将钱i个分配集合为j的情况下的不公平程度最小值
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, 1<<n)
	}
	dp[0] = sum
	for i := 1; i < k; i ++ {
		for j := 0; j < 1<<n; j ++ {
			dp[i][j] = 1e9
			//枚举j的子集s
			for s := j; s > 0 ; s = (s-1) & j {
				dp[i][j] = min(dp[i][j], max(dp[i-1][j^s], sum[s]))
			}
		}
	}
	return dp[k-1][(1<<n)-1]
}
