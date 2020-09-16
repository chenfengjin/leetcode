#
# @lc app=leetcode id=224 lang=python3
#
# [224] Basic Calculator
#

# @lc code=start
import sys


class Solution:
    def calculate(self, s: str) -> int:
        s = "(" + s + ")"
        total = []
        left = 0
        right = 0
        current  = 0
        sign = 1
        while right + 1 < len(s): # right 需要前瞻一个字符
            pass

if __name__ == "__main__":
    print(Solution().calculate("2-(5-6)"))

# @lc code=end
