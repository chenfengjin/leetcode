#
# @lc app=leetcode.cn id=31 lang=python3
#
# [31] 下一个排列
#
from typing import *

# @lc code=start
class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) == 1:
            return nums
        index = len(nums)-1
        while index > 0 and nums[index] <= nums[index-1]:
            index  -= 1
        if index == len(nums) - 1:
            nums[index],nums[index-1] = nums[index-1],nums[index]
            return nums
        nums[index:len(nums)] = list(reversed(nums[index:]))
        # if index == 0:
            # return nums

        # while index < len(nums) -1  and nums[index] < nums[index+1]:
        #     nums[index],nums[index+1] = nums[index+1],nums[index]
        return nums
             
# @lc code=end


if __name__ == "__main__":
    print(Solution().nextPermutation([1,2,3]))