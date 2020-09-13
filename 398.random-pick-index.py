#
# @lc app=leetcode id=398 lang=python3
#
# [398] Random Pick Index
#

# @lc code=start
import random
class Solution:

    def __init__(self, nums: List[int]):
        self.nums = nums

    def pick(self, target: int) -> int:
        index = self.nums.index(target)
        if index == len(self.nums)-1:
            return index
        for i,num in enumerate(self.nums[index+1:]):
            if num == target and random.randint(0,1) == 1:
                index = i
        return index        


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.pick(target)
# @lc code=end

