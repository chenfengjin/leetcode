#
# @lc app=leetcode id=78 lang=python3
#
# [78] Subsets
#

from typing import *
# @lc code=start
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        output = [[]]
        for num in nums:
            output += [curr+[num] for curr in output]
        return output
        
# @lc code=end
if __name__ == "__main__":
    print(Solution().subsets([1,2,3]))

