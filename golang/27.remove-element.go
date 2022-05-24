/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] Remove Element
 */
package main 
import "fmt"

// @lc code=start
func removeElement(nums []int, val int) int {
	left:=0
	right:=0
	for right<len(nums){
		if nums[right]==val{
			right++
			continue	
		}
		nums[left] = nums[right]
		left++
		right++
	}
	return left
}
// @lc code=end

func main(){
	fmt.Println(removeElement([]int{3,2,2,3},3))
}
