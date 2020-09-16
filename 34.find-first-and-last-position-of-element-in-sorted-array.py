#
# @lc app=leetcode id=34 lang=python3
#
# [34] Find First and Last Position of Element in Sorted Array
#
from typing import *
# @lc code=start
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        try:
            return [nums.index(target),len(nums) - list(reversed(nums)).index(target)-1 ]
        except:
            return [-1,-1]
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().searchRange([5,7,7,8,8,10],8))