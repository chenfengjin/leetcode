#
# @lc app=leetcode id=268 lang=python3
#
# [268] Missing Number
#

# @lc code=start
from typing import *
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        bucket = list(range(len(nums)+1))
        for num in nums:
            bucket[num] = -1
        for num in bucket:
            if not num == -1:
                return num
        
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().missingNumber([1,4,5,6,3]))