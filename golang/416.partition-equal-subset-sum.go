/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */
package main

// @lc code=start
func sum(items []int) int {
	total := 0
	for i := 0; i < len(items); i++ {
		total += items[i]
	}
	return total

}
func canPartition(nums []int) bool {
	total := sum(nums)
	if total&1 == 1 {
		return false
	}
	found := false
	collect := func() {
		found = true
	}
	cur := []int{}
	canPartitionHelper(nums, cur, 0, 0, total/2, collect)
	return found
}
func canPartitionHelper(nums, cur []int, curSum, K, target int, collect func()) {
	if curSum == target {
		collect()
		return
	}

	for i := K; i < len(nums); i++ {
		num := nums[i]
		cur = append(cur, num)
		canPartitionHelper(nums, cur, curSum+num, i+1, target, collect)
		cur = cur[:len(cur)-1]
	}
}

// @lc code=end

func main() {
	canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97})
}
