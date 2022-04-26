/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */
// TODO 有一个基于位运算的解决方案
// https://leetcode-cn.com/problems/coin-change/solution/ji-yu-wei-yun-suan-de-dong-tai-gui-hua-p-9m8o/
// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := []int{0}
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && min > dp[i-coins[j]] {
				min = dp[i-coins[j]] + 1
			}
		}
		dp = append(dp, min)
	}
	if dp[len(dp)-1] == math.MaxInt32 {
		return -1
	}
	return dp[len(dp)-1]
}

// @lc code=end

