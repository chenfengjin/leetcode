#
# @lc app=leetcode id=509 lang=python3
#
# [509] Fibonacci Number
#

# @lc code=start
class Solution:
    def fib(self, N: int) -> int:
        if N == 0:
            return 0
        res = [0,1]
        for i in range(2,N+1):
            res.append(res[-1]+res[-2])
        return res[-1]
        
# @lc code=end

