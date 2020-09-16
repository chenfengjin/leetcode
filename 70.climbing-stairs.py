#
# @lc app=leetcode id=70 lang=python3
#
# [70] Climbing Stairs
#

# @lc code=start
class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 0:
            return 1
        if n == 1:
            return 1
        step = [1,1]
        for i in range(2,n+1):
            # print(step,i)
            step.append(step[i-1] + step[i-2])
        # print(step)
        return step[-1]
        
# @lc code=end

