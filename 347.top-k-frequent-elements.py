#
# @lc app=leetcode id=347 lang=python3
#
# [347] Top K Frequent Elements
#
import heapq
from typing import *
# @lc code=start
class Solution:
    # Bad solution
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counts_map = {}
        for num in nums:
            if not num in counts_map:
                counts_map[num] = 0
            counts_map[num] += 1
        values = sorted(counts_map.values(),reverse=True)[:k]
        return [key for (key,value) in counts_map.items() if value in values]
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().topKFrequent([3,0,1,0],1))
    # print(Solution().topKFrequent([0],0))