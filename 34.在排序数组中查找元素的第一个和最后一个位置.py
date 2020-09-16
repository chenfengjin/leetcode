#
# @lc app=leetcode.cn id=34 lang=python3
#
# [34] 在排序数组中查找元素的第一个和最后一个位置
#
from typing import *
# @lc code=start
# 几个需要注意的点
# nums 为空的情况
# 找不到的情况
# 
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        if not nums:
            return [-1,-1]
        left = 0
        right = len(nums)
        while left < right:
            middle = (left + right ) >> 1
            if nums[middle] < target:
                left = middle + 1 
            else:
                right = middle
        if left >= len(nums) or not nums[left] ==  target :
            return [-1,-1]
        while left-1 >= 0 and nums[left-1] == target:
            left -= 1
        while right+1 < len(nums) and nums[right+1] ==  target:
            right += 1
        return [left,right]
                        
# @lc code=end

if __name__ == "__main__":
    print(Solution().searchRange([5,7,7,8,8,10],8))