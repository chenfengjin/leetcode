#
# @lc app=leetcode id=300 lang=python3
#
# [300] Longest Increasing Subsequence
#

# @lc code=start
from typing import *
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        if not nums:
            return 0
        n = len(nums)
        max_lengthes = [1] * n
        i = -1
        for i in range(n):
            for j in range(i-1,-1,-1):
                if nums[i] > nums[j] and max_lengthes[i] < 1+ max_lengthes[j]:
                    max_lengthes[i] = 1+ max_lengthes[j]                    
        print(max_lengthes)
        return max(max_lengthes)

        
# @lc code=end



if __name__ == "__main__":
    print(Solution().lengthOfLIS([10,9,2,5,3,7,101,18]))