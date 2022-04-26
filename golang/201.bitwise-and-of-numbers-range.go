/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

package main
import(
	"fmt"
)
// pay attention to start and limit of i
// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	i:=31
	mask:=0
    for ;i>=0;i--{
		if (left>>i &1) ^ (right>>i&1) !=0{
			break	
		}
		mask |=1<<i
	}
	return right &mask
}
// @lc code=end

func main(){
	fmt.Println(rangeBitwiseAnd(5,7))
}