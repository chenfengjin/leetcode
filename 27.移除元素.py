#
# @lc app=leetcode.cn id=27 lang=python3
#
# [27] 移除元素
#

# @lc code=start
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        left = 0
        right  = 0
        if not nums:
            return 0
        while right < len(nums):
            if not nums[right] == val:
                left += 1
                nums[left] = nums[right]
            right += 1
        return left+1
# @lc code=end

