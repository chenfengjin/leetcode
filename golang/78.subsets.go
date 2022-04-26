/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
//  TODO this is only for nums with no more 32 elements on 32 bit machine and 64 for 64 bit machine
//  use int64 for n may be more portable for golang, but benifit is limited
// fortunately, all test cases pass for leetcode
package main

import "fmt"

// @lc code=start
func subsets(nums []int) [][]int {
	n := 1 << len(nums)
	ret := [][]int{}
	for i := 0; i < n; i++ {
		item := []int{}
		m := 0
		j := i
		for j > 0 {
			if j&1 == 1 {
				item = append(item, nums[m])
			}
			j >>= 1
			m++
		}
		ret = append(ret, item)

	}
	return ret

}

// @lc code=end
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
