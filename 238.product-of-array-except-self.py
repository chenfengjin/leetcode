#
# @lc app=leetcode id=238 lang=python3
#
# [238] Product of Array Except Self
#

# @lc code=start
from typing import *
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        left = [1]
        right = [1]
        for index,num in enumerate(nums):
            left.append(num*left[index])
        for index,num in enumerate(list(reversed(nums))):
            right.append(right[index]*num)
        return [left[i]*right[len(nums)-i-1] for i in range(len(nums))]
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().productExceptSelf([1,2,3,4]))