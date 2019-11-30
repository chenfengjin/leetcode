/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 *
 * https://leetcode.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (40.12%)
 * Total Accepted:    110.1K
 * Total Submissions: 274.3K
 * Testcase Example:  '16'
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a
 * power of 4.
 * 
 * Example 1:
 * 
 * 
 * Input: 16
 * Output: true
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 5
 * Output: false
 * 
 * 
 * Follow up: Could you solve it without loops/recursion?
 */
func isPowerOfFour(num int) bool {
	return (num==1<<0)||num==1<<2||(num == 1<<4)||(num == 1<<6)||(num == 1<<8)||(num == 1<<10)||(num == 1<<12)||(num == 1<<14)||(num == 1<<16)||(num == 1<<18)||(num == 1<<20)||(num == 1<<22)||(num == 1<<24)||(num == 1<<26)||(num == 1<<28)||(num == 1<<30)
}

