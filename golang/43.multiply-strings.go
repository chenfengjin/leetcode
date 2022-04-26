/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */
package main
// @lc code=start
func multiply(num1 string, num2 string) string {
	suffix := ""
	all:=""
	for i:=len(num2)-1;i>=0;i--{
		one:=multiplyOne(num1,num2[i])
		all  = add(all,one+suffix)
		suffix = suffix+"0"	
	}
	return all
    
}
func multiplyOne(num1,num2 string )string{

}

func add(num1,num2 string)string{

}
// @lc code=end

// all=0
// all=all+234
// all=all+4680
// all=all+70200 
// 	234
// 	321
// 	234
// 4680
