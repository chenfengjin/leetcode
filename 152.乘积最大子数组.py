#
# @lc app=leetcode.cn id=152 lang=python3
#
# [152] 乘积最大子数组
#
from typing import *
# @lc code=start
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        if not nums:
            return 0
        max_factory = nums[0]
        min_factory = nums[0]
        res = nums[0]
        for num in nums:
            # bug point
            max_factory = max(max_factory*num,min_factory*num,num)
            min_factory = min(max_factory*num,min_factory*num,num)
            res = max_factory
        return res
# @lc code=end

if __name__ == "__main__":
    print(Solution().maxProduct([2,3,-2,4]))