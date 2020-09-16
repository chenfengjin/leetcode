/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (59.51%)
 * Total Accepted:    441.2K
 * Total Submissions: 740.4K
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 * 
 * Note:
 * 
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 * 
 * Example 1:
 * 
 * 
 * Input: [2,2,1]
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,1,2,1,2]
 * Output: 4
 * 
 * 
 */
func singleNumber(nums []int) int {

	for _,num:=range nums[1:]{
		nums[0]=nums[0] ^ num
	}
	return nums[0]
}

