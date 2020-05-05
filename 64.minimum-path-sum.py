#
# @lc app=leetcode id=64 lang=python3
#
# [64] Minimum Path Sum
#

# @lc code=start
class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        dp = [[0 for i in grid[0]]  for j in grid]
        for i,row in enumerate(grid):
            for j,cell in enumerate(row):
                if i == 0 and j == 0:
                    dp[i][j] = cell
                    continue
                if i == 0:
                    dp[i][j] = dp[i][j-1] + cell
                    continue
                if j == 0:
                    dp[i][j] =  dp[i-1][j] + cell
                    continue
                dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + cell
        print(i,j,dp)
        return dp[i][j]                        
# @lc code=end

