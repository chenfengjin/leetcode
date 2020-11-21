#
# @lc app=leetcode.cn id=41 lang=python3
#
# [41] 缺失的第一个正数
#
from typing import *
# @lc code=start
class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        for index,num in enumerate(nums):
            while 0 < index < len(nums) and not nums[index] == index:
                tmp = nums[index]
                nums[index] = num
                index = tmp
        # print(nums)
        for index,num in enumerate(nums[1:],start=1):
            if not index == num:
                return index
        return len(nums)
# @lc code=end

if __name__ == "__main__":
    # print(Solution().firstMissingPositive([3,4,-1,1]))
    print(Solution().firstMissingPositive([1,2,0]))