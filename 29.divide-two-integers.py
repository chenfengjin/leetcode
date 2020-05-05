#
# @lc app=leetcode id=29 lang=python3
#
# [29] Divide Two Integers
#

# @lc code=start
class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        sign = 1 if dividend > 0 == divisor > 0 else -1
        dividend = abs(dividend)
        divisor = abs(divisor)
        summary = divisor
        count = 0
        while summary + divisor < dividend:
            count += 1
            summary += divisor
        return sign * count

        
# @lc code=end

