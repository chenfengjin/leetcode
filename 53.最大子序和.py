#
# @lc app=leetcode.cn id=53 lang=python3
#
# [53] 最大子序和
#
from typing import *
# @lc code=start
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if not nums:
            return 0
        dp = [nums[0]]
        for num in nums[1:]:
            dp.append(num+dp[-1] if dp[-1] >0 else num)
        return max(dp)
    # TODO 试一下分治？
# @lc code=end
if __name__ == "__main__":
    print(Solution().maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))
