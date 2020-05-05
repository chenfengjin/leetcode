#
# @lc app=leetcode id=378 lang=python3
#
# [378] Kth Smallest Element in a Sorted Matrix
#

# @lc code=start
class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        # ugly answer
        l = []
        for i in matrix:
            l.extend(i)
        return sorted(l)[k-1]       
# @lc code=end

