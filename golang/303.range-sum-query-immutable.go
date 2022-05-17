/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */
package main

// @lc code=start
type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	total := 0
	prefixSum := []int{0}
	for _, num := range nums {
		total += num
		prefixSum = append(prefixSum, total)
	}
	return NumArray{
		prefixSum: prefixSum,
	}

}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
