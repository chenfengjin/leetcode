#
# @lc app=leetcode id=650 lang=python3
#
# [650] 2 Keys Keyboard
#

# @lc code=start
from typing import *
class Solution:
    def minSteps(self, n: int) -> int:
        dp=[0,0]
        for i in range(2,n):
            min_step = 1000 
            for j in range(1,i):
                if i % j == 0:
                    min_step = min(min_step,dp[j]+i//j)
            dp.append(min_step)
        return dp[n-1]
                
        

        
# @lc code=end


if __name__ == "__main__":
    print(Solution().minSteps(3))