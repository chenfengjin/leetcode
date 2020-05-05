#
# @lc app=leetcode id=198 lang=python3
#
# [198] House Robber
#

# @lc code=start
class Solution: #动态规划千万不要只比较最后两个数
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

