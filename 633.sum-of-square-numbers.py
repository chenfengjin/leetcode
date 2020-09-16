#
# @lc app=leetcode id=633 lang=python3
#
# [633] Sum of Square Numbers
#

# @lc code=start
class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        if c == 0:
            return True
        for i in range(1,math.sqrt(c)):
            if round(math.sqrt(c-i)) ** 2 == c-i:
                return True
        return False
# @lc code=end

