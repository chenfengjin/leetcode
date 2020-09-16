#
# @lc app=leetcode id=55 lang=python3
#
# [55] Jump Game
#

# @lc code=start
class Solution:
    def canJump(self, nums: List[int]) -> bool:
        reach = 0
        i = 0
        
        while i < len(nums):
            if i > reach:
                return False
            reach = max(reach , i + nums[i] )
            i = i + 1
        return True  

# @lc code=end

