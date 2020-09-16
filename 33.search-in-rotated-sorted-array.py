#
# @lc app=leetcode id=33 lang=python3
#
# [33] Search in Rotated Sorted Array
#

from typing import *
# @lc code=start
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if not nums:
            return -1
        if len(nums) == 1:
            if nums[0] == target:
                return 0
            return -1

        if nums[0] < nums[-1]:
            return self.binary_search(nums,target)

        for i in range(len(nums)):
            if nums[i] > nums[i+1]:
                break
        if target >= nums[0]:  
            return self.binary_search(nums[0:i+1],target)

        sub_index = self.binary_search(nums[i+1:],target)
        return -1 if sub_index == -1 else sub_index + i +1

    def binary_search(self,nums,target):
        left = 0
        right = len(nums) - 1
        while right - left > 1:
            middle = (left + right )//2
            if nums[middle] == target:
                return middle
            if target > nums[middle]:
                left = middle
            if target < nums[middle]:
                right = middle
        if nums[left] == target:
            return left
        if nums[right] == target:
            return right
        return -1

            
# @lc code=end


if __name__ == "__main__":
    print(Solution().binary_search([1,2,3,5,6,7,8],4))
    print(Solution().binary_search([1,2,3,5,6,7,8],4))
    print(Solution().binary_search([1,2,3,5,6,7,9],5))
    print(Solution().binary_search([1,2,3,5,6,7,9],5))