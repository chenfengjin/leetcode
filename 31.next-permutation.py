#
# @lc app=leetcode id=31 lang=python3
#
# [31] Next Permutation
#
from typing import *
# @lc code=start
class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        left = len(nums) -1 -1 
        right = len(nums)-1
        while left >=0 and nums[left] >= nums[left+1] : # -1 is also valid index in python
            left -= 1
        if left == -1:
            if sorted(nums,reverse=True) == nums:
                nums.sort()
                return
            else:
                left = 0
        while right >= 0 and nums[left] >= nums[right]:
            right -= 1
        nums[left],nums[right] = nums[right],nums[left]
        nums[left+1:] = sorted(nums[left+1:])

# @lc code=end


if __name__ == "__main__":
    nums = [5,1,1,]
    Solution().nextPermutation(nums)
    print(nums)