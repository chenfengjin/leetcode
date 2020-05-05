#
# @lc app=leetcode id=416 lang=python3
#
# [416] Partition Equal Subset Sum
#
from typing import *
# @lc code=start
class Solution:
    def helper(self,nums,k):
        if k in nums:
            return True
        for  (index,num) in enumerate(nums):
            if num > k:
                continue
            if self.helper(nums[index+1:],k-num):
                return True
        return False

    def canPartition(self, nums: List[int]) -> bool:
        sum_of_all = sum(nums)
        if sum_of_all % 2 == 1:
            return False
        return self.helper(nums,sum_of_all//2)            
            
        
# @lc code=end

if __name__ == "__main__":
    a=[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,100]
    print(Solution().canPartition(a))
