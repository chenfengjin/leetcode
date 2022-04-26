/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */
 package main
 import "fmt"
//  you are not asked to return all 132 pattern
//  so stack is neccessary
// @lc code=start
func find132pattern(nums []int) bool {
	length := len(nums)
	l1 := make([]int, length)
	l2 := make([]int, length)
	l1[0] = nums[0]
	l2[0] = nums[length-1]
	for i := 1; i < len(nums); i++ {
		if nums[i] < l1[i-1] {
			l1[i] = nums[i]
		} else {
			l1[i] = l1[i-1]
		}
		if nums[length-i-1] < l2[i-1]{
			l2[i] = nums[length-i-1]
		}else{
			l2[i] = l2[i-1]
		}
	}
	fmt.Println(l1)
	fmt.Println(l2)
	for i:=1;i<length-1;i++{
		if nums[i] > l1[i] && nums[i] > l2[length-1-i]{
			return true
		}
	}
	return false
}

// @lc code=end


func main(){
	fmt.Println(find132pattern([]int{1,0,1,-4,-3}))
}