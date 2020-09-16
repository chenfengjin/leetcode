#
# @lc app=leetcode id=275 lang=python3
#
# [275] H-Index II
#

# @lc code=start
class Solution:
    def hIndex(self, citations: List[int]) -> int:
        for i in range(1,len(citations) + 1):
            if citations[-1*i] <i:
                return i-1
        return len(citations) if len(citations) > 0 and  max(citations) > 0 else 0
# @lc code=end

