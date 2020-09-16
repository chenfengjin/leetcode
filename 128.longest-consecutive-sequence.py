#
# @lc app=leetcode id=128 lang=python3
#
# [128] Longest Consecutive Sequence
#

from typing import *
# @lc code=start
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not nums:
            return 0
        m = {}
        for num in nums:
            if num in m:
                continue
            if num-1 in m:
                if num+1 in m:
                    m[num] = m[num-1] + m[num+1] +1
                    m[num-1] = m[num]
                    m[num+1] = m[num]
                else:
                    m[num] = m[num-1]+1
                    m[num-1] = m[num]
            elif num+1 in m:
                m[num] = m[num+1] + 1
                m[num+1] += 1
            else:
                m[num] = 1
        return max(m.values())
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().longestConsecutive([0,3,7,2,5,8,4,6,0,1]))