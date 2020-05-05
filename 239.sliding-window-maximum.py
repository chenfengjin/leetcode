#
# @lc app=leetcode id=239 lang=python3
#
# [239] Sliding Window Maximum
#

# @lc code=start
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]: #bad answer
        if not nums:
            return []
        res = []
        for i in range(len(nums)-k+1):
            res.append(max(nums[i:i+k]))
        return res
# @lc code=end

