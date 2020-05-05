#
# @lc app=leetcode id=347 lang=python3
#
# [347] Top K Frequent Elements
#
import heapq
from typing import *
# @lc code=start
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counts_map = {}
        for num in nums:
            if not num in counts_map:
                counts_map[num] = 0
            counts_map[num] += 1
        counts = list(counts_map.values())
        reversed_counts = []
        heapq.heapify(counts)
        while counts:
            reversed_counts.append(heapq.heappop(counts))
        print(reversed_counts[k-1])
        return [i for i in counts_map if not counts_map[i] < reversed_counts[k-1]]
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().topKFrequent([3,0,1,0],1))
    # print(Solution().topKFrequent([0],0))