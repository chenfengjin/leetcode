/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (53.95%)
 * Total Accepted:    210.1K
 * Total Submissions: 389.4K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 * 
 * Example 1:
 * 
 * 
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 * 
 * 
 * Note:
 * 
 * 
 * Each element in the result must be unique.
 * The result can be in any order.
 * 
 * 
 * 
 * 
 */
func intersection(nums1 []int, nums2 []int) []int {
	m:=map[int]int{}
	result:=[]int{}

	for _,num:=range nums1{
		m[num]=1
	}

	for _,num:=range nums2{
		if _,ok:=m[num];ok{
			m[num]=2
		}
	}

    for k,v:=range m{
		if v==2{
			result=append(result,k)
		}
	}
	return result
}

