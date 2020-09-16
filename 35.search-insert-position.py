#
# @lc app=leetcode id=35 lang=python3
#
# [35] Search Insert Position
#
from typing import *
# @lc code=start
class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left = 0
        right = len(nums)
        while left < right:
            middle = (left+right) // 2
            if nums[middle] < target:
                left = middle + 1
            else:
                right = middle
        return left
if __name__ == "__main__":
    print("\tleft\tright\tmiddle")
    print("for zero:")
    # print(Solution().searchInsert([1,3,5,6],0))
    print("for two:")
    print(Solution().searchInsert([1,3,5,6],5))
# @lc code=end

