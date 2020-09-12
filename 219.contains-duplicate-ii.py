#
# @lc app=leetcode id=219 lang=python3
#
# [219] Contains Duplicate II
#

# @lc code=start
from typing import * 

class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        nums_table = {}
        for index,num in enumerate(nums):
            if num in nums_table and index - nums_table.get(num) <= k:
                return True
            nums_table[num] = index
        return False
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().containsNearbyDuplicate([1,2,3,1],3))