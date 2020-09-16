#
# @lc app=leetcode id=45 lang=python3
#
# [45] Jump Game II
#

# @lc code=start
class Solution:
    def jump(self, nums: List[int]) -> int:
        reach = 0
        i = 0
        steps = 0
        
        while i < len(nums):
            if i > reach:
                return False
            reach = max(reach , i + nums[i] )
            i = i + 1
        return True  
# @lc code=end

