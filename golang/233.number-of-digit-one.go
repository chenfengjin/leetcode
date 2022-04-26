/*
 * @lc app=leetcode id=233 lang=golang
 *
 * [233] Number of Digit One
 */
package main

//  digit dynamic programming is naturally suitable for this problem
// @lc code=start
var (
	digits []int{}
	dp 	[]int{}
)
func countDigitOne(n int) int {
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n %= 10
	}
	pos:=0
}

func dfs(digits,pos,dp){
	ret:=0
	for i:=0;i<len(digits);i++{
		for j:=0;j<digits[i];j++{
			ret+=dfs(digits,n+1)
		}
	}

}
// @lc code=end
func main() {
	// fmt.Printf("%d", countDigitOne(13))
}
