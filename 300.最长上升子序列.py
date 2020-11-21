#
# @lc app=leetcode.cn id=300 lang=python3
#
# [300] 最长上升子序列
#
from typing import *
# @lc code=start
class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        if not nums:
            return 0
        dp = [0]
        i = 1
        # TODO 可以对已经搜索过的部分进行排序，然后进行二分查找，可以实现O(nlogn)
        while i < len(nums):
            prev = [z[0] + 1  for z in zip(dp,nums) if z[1] < nums[i]]
            if prev:
                dp.append( max(prev))
            else:
                dp.append(1)
            i += 1 
        return max(dp)
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().lengthOfLIS([10,9,2,5,3,7,101,18]))