#
# @lc app=leetcode id=152 lang=python3
#
# [152] Maximum Product Subarray
#

from typing import *
# @lc code=start
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        if not nums:
            return 0

        max_array = [1]
        min_array = [-1]
        for index, num in enumerate(nums):
            if min_array[-1] == 0:
                max_array.append(num)
                min_array.append(num)
                continue
            if num <0:
                max_array.append(min_array[index-1] * num)
                min_array.append(max_array[index-1] * num)
            else:
                max_array.append(max_array[index-1] * num)
                min_array.append(min_array[index-1] * num)
                
        return max(max_array)
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().maxProduct([3,-1,4]))