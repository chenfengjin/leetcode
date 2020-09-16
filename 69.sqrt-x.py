#
# @lc app=leetcode id=69 lang=python3
#
# [69] Sqrt(x)
#

# @lc code=start
class Solution:
    def mySqrt(self, x: int) -> int:
        left = 0
        right  = x
        while left <= right:
            middle = (left + right) // 2
            square = middle * middle
            if square ==  x:
                return middle
            else:
                if square > x:
                    right = middle - 1
                else:
                    left = middle + 1
        return middle

                        
# @lc code=end

