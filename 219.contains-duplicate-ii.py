#
# @lc app=leetcode id=219 lang=python3
#
# [219] Contains Duplicate II
#

# @lc code=start
from typing import * 
# 用哈希表，而不是集合
class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        nums_table = set()
        for num in nums[0:k]:
            if num in nums_table:
                return True
            nums_table.add(num)

        for index in range(k,len(nums)):
            if nums[index]  in nums_table:
                return True
            nums_table.add(nums[index])
            if nums[index-k] in nums_table:
                nums_table.remove(nums[index-k])
        return False
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().containsNearbyDuplicate([1,2,3,1],3))