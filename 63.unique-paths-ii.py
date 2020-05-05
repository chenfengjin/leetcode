#
# @lc app=leetcode id=63 lang=python3
#
# [63] Unique Paths II
#

# @lc code=start
class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        pass
    def combination(self,n,k): # 坑在浮点预算
        res = 1
        i = k
        while i > 0:
            res *=  n
            n = n-1
            i -= 1
        i = k
        while i > 0 :
            res = res // i
            i -= 1
        return res
    def uniquePaths(self, m: int, n: int) -> int:
        return self.combination(m+n-2,n-1)
        
# @lc code=end

