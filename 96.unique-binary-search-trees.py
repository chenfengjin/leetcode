#
# @lc app=leetcode id=96 lang=python3
#
# [96] Unique Binary Search Trees
#

# @lc code=start
class Solution:
    def numTrees(self, n: int) -> int:
        if n == 1:
            return 1
        res = 0
        for i in range(n):
            res += self.numTrees(i) * self.numTrees(n-i)
        
# @lc code=end

