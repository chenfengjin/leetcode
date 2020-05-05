#
# @lc app=leetcode id=32 lang=python3
#
# [32] Longest Valid Parentheses
#

# @lc code=start
class Solution:
    def __init__(self):
        self.inner_index= []
    def longestValidParentheses(self, s: str) -> int:
        if len(s) < 2:
            return 0
        self.prerocess(s)
        if not self.inner_index:
            return 0
        return max([self.expand(s,i,i+1) for i in self.inner_index])

    def prerocess(self,s):
        for i in range(len(s)-1):
            print(s[i:i+2] )
            if s[i:i+2] == "()":
                self.inner_index.append(i)

    def expand(self,s,left,right):
        print(left,right,s[left:right+1])
        if left < 0 or right > len(s)-1:
            return right - left
        if left > 0 and right < len(s) -1 and  s[left-1] == '(' and s[right+1] == ')':
            return self.expand(s,left-1,right+1)
        return max(self.longestValidParentheses(s[0:left]),self.longestValidParentheses(s[right])) + right-left +1
        
# @lc code=end

if __name__ == "__main__":
    s=")(((((()())()()))()(()))("
    print(Solution().longestValidParentheses(s))
    # print(Solution().expand("((()))())",0,5))