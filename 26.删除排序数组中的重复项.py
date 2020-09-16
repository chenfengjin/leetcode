#
# @lc app=leetcode.cn id=26 lang=python3
#
# [26] 删除排序数组中的重复项
#
from typing import *
# @lc code=start
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        left = 0
        right  = 0
        if not nums:
            return 0
        while right < len(nums):
            if not nums[left] == nums[right]: #TODO  这里解决了下标为 0 的情况
                left += 1
                nums[left] = nums[right]
            right += 1
        return left+1
# @lc code=end

if __name__ == "__main__":
    nums = [1,1,2]
    Solution().removeDuplicates(nums)
    print(nums)