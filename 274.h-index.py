#
# @lc app=leetcode id=274 lang=python3
#
# [274] H-Index
#

# @lc code=start
from typing import *
class Solution:
    def hIndex(self, citations: List[int]) -> int:
        if not citations or max(citations) == 0:
            return 0
        citations.sort(reverse=True)
        h_index = 0
        for index,citation in enumerate(citations):
            if  not citation < index + 1:
                h_index += 1
        return h_index
        


# @lc code=end

if __name__ == "__main__":
    print(Solution().hIndex([3,0,6,1,5]))
