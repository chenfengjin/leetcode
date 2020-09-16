#
# @lc app=leetcode id=704 lang=python3
#
# [704] Binary Search
#

# @lc code=start
class Solution:
    def search(self, nums: List[int], target: int) -> int:
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

