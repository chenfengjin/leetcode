/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (53.96%)
 * Total Accepted:    320.7K
 * Total Submissions: 592.9K
 * Testcase Example:  '3'
 *
 * 
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 * 
 * 
 * 
 * For example, given n = 3, a solution set is:
 * 
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
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

func generateParenthesis(n int) []string {
	paren:=""
}

