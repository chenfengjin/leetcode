#
# @lc app=leetcode id=80 lang=python3
#
# [80] Remove Duplicates from Sorted Array II
#
from typing import *
# @lc code=start
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int: 
        if len(nums) < 3:
            return len(nums)
        left = 2
        right = 2
        buffer = []
        while right < len(nums):
            if nums[right] == nums[right-1] and nums[right] == nums[right-2]: # 原地修改之后这里可能会变化
                right += 1
            else:
                nums[left] = nums[right]
                left += 1
                right += 1

        return left 

# @lc code=end

if __name__ == "__main__":
    print(Solution().removeDuplicates([1,1,1,2,2,3]))