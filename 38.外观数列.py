#
# @lc app=leetcode.cn id=38 lang=python3
#
# [38] 外观数列
#

# @lc code=start
class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"
        count = 0
        index = 0 
        previous = self.countAndSay(n-1)
        while index < len(previous):
            if index and not previous[index]
# @lc code=end

