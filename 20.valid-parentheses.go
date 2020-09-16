/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (36.19%)
 * Total Accepted:    552.7K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */
func isValid(s string) bool {
	if s==""{
		return true
	}
	stackS:=[]byte{}
	for i:=0;i<len(s);i++{
		if s[i]=='{' ||s[i]=='['||s[i]=='('{
			stackS=append(stackS,s[i])
		}else{
			if len(stackS)==0{
				return false
			}
			if (stackS[len(stackS)-1]=='{' && s[i]=='}')||(stackS[len(stackS)-1]=='[' &&s[i]==']')||(stackS[len(stackS)-1]=='(' && s[i]==')'){
				stackS = stackS[0:len(stackS)-1]	
			}else{
				return false
			}
		}
	}
	if len(stackS)==0{
		return true
	}
	return false
}

