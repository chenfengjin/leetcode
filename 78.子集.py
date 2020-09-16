#
# @lc app=leetcode.cn id=78 lang=python3
#
# [78] 子集
#
from typing import *
# @lc code=start
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        current = []
        self.helper(result,current=current,nums = nums,index = 0)
        return result

    def helper(self,result,current,nums,index):
        result.append(current)
        for i in range(index,len(nums)):
            self.helper(result,current+[nums[i]],nums,i + 1)

# @lc code=end

if __name__ == "__main__":
    print(Solution().subsets([1,2,3]))