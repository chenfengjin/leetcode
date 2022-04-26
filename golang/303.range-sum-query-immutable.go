/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */
package main

// @lc code=start
type NumArray1 struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	preSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		preSum = append(preSum, preSum[i-1]+nums[i])
	}
	return NumArray{
		preSum: preSum,
	}

}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.preSum[right]
	}
	return this.preSum[right] - this.preSum[left-1]

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
