#
# @lc app=leetcode id=16 lang=python3
#
# [16] 3Sum Closest
#
from typing import *
# @lc code=start
class Solution:
    # 通过双指针，把 O(n^3) 的复杂度降低到 O(n^2)
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        min_gap = abs(sum(nums[0:3])-target)
        result = sum(nums[0:3])

        for index,num in enumerate(nums):
            left = 0
            right = len(nums)-1
            remain = target-nums[index]
            while left < index and right > index : # Do not use left < index # 这里可能有错？
                gap = abs(nums[left] + nums[right] - remain) 
                if gap < min_gap:
                    min_gap = gap
                    result =  nums[left] + nums[right] + nums[index]
                if nums[left] + nums[right] > remain :
                    right -= 1
                elif nums[left] + nums[right] < remain:
                    left += 1 
                else: # 这里别丢
                    return target
        return result
# @lc code=end

if __name__ == "__main__":
    print(Solution().threeSumClosest([-1, 2, 1, -4],1))