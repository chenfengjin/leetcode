#
# @lc app=leetcode id=84 lang=python3
#
# [84] Largest Rectangle in Histogram
#
from typing import *
# @lc code=start
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        if not heights:
            return 0

        return max([heights[index]*self.expand(heights,index) for index in range(len(heights))])
    def expand(self,heights,index):
        left = index
        right = index + 1
        while left > 0 and heights[left-1] >= heights[index]:
            left -= 1
        while right  < len(heights) and heights[right] >= heights[index]:
            right += 1
        
        return right - left 
        
# @lc code=end

if __name__ == "__main__":
    # print(Solution().expand([2,1,5,6,2,3],2))
    print(Solution().largestRectangleArea([]))
    print(Solution().largestRectangleArea([1]))
    print(Solution().largestRectangleArea([1,2]))
