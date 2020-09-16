#
# @lc app=leetcode id=228 lang=python3
#
# [228] Summary Ranges
#

# @lc code=start
class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if not nums:
            return []
        if len(nums) == 1:
            return [str(nums[0])]
        res = []
        left = 0
        right = 1
        while right <= len(nums):
            if right == len(nums) or nums[right] - nums[right-1] > 1:
                if left == right - 1:
                    res.append(str(nums[left]))
                else:
                    res.append("{}->{}".format(nums[left],nums[right-1]))
                left = right
                
            right = right + 1
        return res
        
# @lc code=end

