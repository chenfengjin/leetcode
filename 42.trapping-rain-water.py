#
# @lc app=leetcode id=42 lang=python3
#
# [42] Trapping Rain Water
#

from typing import *
# @lc code=start
class Solution:
    def trap(self, height: List[int]) -> int:
        max_left = [0] * len(height)
        max_right = [0] * len(height)
        max_value = 0
        for index,h in enumerate(height):
            if h > max_value:
                max_value = h
            max_left[index] = max_value
        max_value = 0

        for index,h in enumerate(reversed(height)):
            if h> max_value:
                max_value = h
            max_right[index] = max_value
        max_right.reverse()
        total = []
        for index in range(len(height)):
            if  min(max_left[index],max_right[index]) > height[index]:
                total.append(min(max_left[index],max_right[index]) - height[index])
            else:
                total.append(0)
        return sum(total)
                 
        print(max_left)
# @lc code=end

if __name__ == "__main__":
    print(Solution().trap([0,1,0,2,1,0,1,3,2,1,2,1]))