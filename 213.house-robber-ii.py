#
# @lc app=leetcode id=213 lang=python3
#
# [213] House Robber II
#

# @lc code=start
class Solution:
    def rob(self, nums: List[int]) -> int:
        if not nums:
            return 0
        if len(nums) <= 2:
            return max([num for num in nums])
        dp = [nums[0]]
        dp.append(nums[1] if nums[1] > nums[0] else nums[0]) 
        for i in range(2,len(nums)):
            if nums[i] + dp[i-2] > dp[i-1]:
                dp.append(dp[i-2] + nums[i])
            else:
                dp.append(dp[i-1])     
        return max(dp)
        
# @lc code=end

