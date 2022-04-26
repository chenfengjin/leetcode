/*
 * @lc app=leetcode id=600 lang=golang
 *
 * [600] Non-negative Integers without Consecutive Ones
 */
package main
import (
	"fmt"
)
// @lc code=start
func findIntegers(n int) int {
	ret:=0
	for i:=1;i<n;i++{
		if !findIntegersOne(i){
			ret+=1
		}
	}
	return ret
}
func findIntegersOne(n int)bool{
	mask1:=0x33333333
	mask2:=0x66666666
	mask3:=0xCCCCCCCC
	mask4:=0x18181818
	return (n&mask1!=0) || (n&mask2!=0)|| (n&mask3!=0)||(n&mask4!=0)
}
// @lc code=end
func main(){
	for i:=0;i<20;i++{
		fmt.Println(i,findIntegersOne(i))
	}
}
