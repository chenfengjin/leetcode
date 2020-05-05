#
# @lc app=leetcode id=224 lang=python3
#
# [224] Basic Calculator
#

# @lc code=start
class Solution:
    def calculate(self, s: str) -> int:
        left = 0
        right = 0 
        stack = []
        sign = 1
        res = 0
        while left < right:
            if s[right] == "(":
                stack.append(res)
            elif s[right] == ")"


        
# @lc code=end

