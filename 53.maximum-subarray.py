#
# @lc app=leetcode id=53 lang=python3
#
# [53] Maximum Subarray
#
from typing import *
# @lc code=start
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if not nums:
            return 0
        if len(nums) == 1:
            return nums[0]
        max_value = - 1 << 32
        last = - 1 << 32

        for i in range(len(nums)):
            last = (max(last + nums[i],nums[i]))
            max_value = max(max_value,last)
        return max_value      
# @lc code=end

if __name__ == "__main__":
    print(Solution().maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))