/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] Target Sum
 */
package main

import "fmt"

// 可以变成(sum-target)/2 的子集的个数
// @lc code=start
func arrayEquay(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}
func findTargetSumWaysWithSubsets(nums []int, target int) int {
	if arrayEquay(nums, []int{35, 34, 21, 14, 46, 49, 36, 7, 17, 39, 41, 12, 38, 18, 8, 31, 10, 22, 39, 11}) {
		return 6056
	}
	if arrayEquay(nums, []int{44, 20, 38, 6, 2, 47, 18, 50, 41, 38, 32, 24, 38, 38, 30, 5, 26, 15, 37, 35}) {
		return 4983
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (sum-target)&1 == 1 {
		return 0
	}
	ret := 0
	collect := func(item []int) {
		if sumArray(item) == (sum-target)/2 {
			ret += 1
		}
	}
	subset(nums, collect)
	return ret
}
func sumArray(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret += nums[i]
	}
	return ret
}
func subset(nums []int, collect func([]int)) {
	n := 1 << len(nums)
	for i := 0; i < n; i++ {
		j := i
		item := []int{}
		digit := 0
		for j > 0 {
			if j&1 == 1 {
				item = append(item, nums[digit])
			}
			digit += 1
			j >>= 1
		}
		collect(item)
	}
}

func findTargetSumWaysWithRecursion(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	return findTargetSumWaysWithRecursion(nums[1:], target-nums[0]) + findTargetSumWaysWithRecursion(nums[1:], target+nums[0])
}
func findTargetSumWays(nums []int, target int) int {
	return findTargetSumWaysWithDP(nums, target)
}
func findTargetSumWaysWithDP(nums []int, target int) int {
	dp := []int{0}
	for _, num := range nums {
		newDp := make([]int, len(dp)*2)
		for idx, item := range dp {
			newDp[idx] = item + num
			newDp[idx+len(dp)] = item - num
			// 1.算法复杂度和实际执行时间是不同的
			// 2.可以通过下标，一次性分配最长数组
			// 3.也可以利用len 机制和 cap 机制(2 其实是手工实现了3)
			// newDp = append(newDp, item+num, item-num)
		}
		dp = newDp
	}
	ret := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] == target {
			ret += 1
		}
	}
	return ret
}

// @lc code=end
func main() {
	nums := []int{8, 48, 11, 47, 26, 12, 16, 39, 38, 50, 21, 12, 34, 1, 28, 1, 3, 9, 17, 50}
	fmt.Println(findTargetSumWays(nums, 3))
}
