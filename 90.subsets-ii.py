#
# @lc app=leetcode id=90 lang=python3
#
# [90] Subsets II
#
from typing import *
# @lc code=start
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        output = [[]]
        left = 0
        for index,num in enumerate(nums):
            if index >0 and num == nums[index-1]:
                output += [curr+[num] for curr in output if not   len(curr) < index - left  and 
                num in curr[left-index:] and len(set(curr[left-index:])) == 1 ]
            else:
                output += [curr+[num] for curr in output]
                left = index
        return output
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().subsetsWithDup([5,5,5,5,5]))