#
# @lc app=leetcode id=164 lang=python3
#
# [164] Maximum Gap
#

# @lc code=start
# 桶排序
class Solution:
    def maximumGap(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return 0
        nums.sort()
        # nums.insert(0,0)
        differences= []
        index = 0
        while index +1 < len(nums):
            differences.append(nums[index+1]-nums[index])
            index += 1
        return max(differences)

        
# @lc code=end

