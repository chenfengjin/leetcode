/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
package main
import "fmt"
// pay attention to conner case 
// stop of j should be nums[i] rather than nums[i]-1
// using negative to take advantage of initializatin of slice 
// @lc code=start
func jump(nums []int) int {
	// 2,3,1,1,4
	// 0 1 1 2 2
	steps := make([]int, len(nums))
	// fmt.Println(len(steps))
	steps[0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		for j := 1; j <=nums[i]; j++ {
			// small trick to use initialization of slice
			if i+j< len(steps) && -1*steps[i+j] >  -1*(nums[i]+1){
				steps[i+j] = -1* (nums[i]+1)
			}
		}
	}
	return steps[len(steps)-1]*-1
}

// @lc code=end

func main(){
	fmt.Println(jump([]int{2,3,1,1,4}))
}