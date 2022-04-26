/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */
package main
// @lc code=start
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
//  func guess(num int) int
func guessNumber(n int) int {
	left:=0
	right:=n
	var middle int 

	for left < right{
		middle=(left+right)/2
		cur:=guess(middle)
		// fmt.Println(cur,left,right,middle)
		if cur ==0{
			return middle
		}else if cur ==1 {
			left = middle+1
		}else{
			right = middle
		}
	}
    return left
}
// @lc code=end

func main(){

}