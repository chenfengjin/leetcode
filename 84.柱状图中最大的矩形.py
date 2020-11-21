#
# @lc app=leetcode.cn id=84 lang=python3
#
# [84] 柱状图中最大的矩形
#
from typing import *
# @lc code=start
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        heights = [0]  + heights + [0]
        stack = []
        largest = 0
        for index,height in enumerate(heights):
            while stack and heights[stack[-1]] > height:
                prev = stack[-1]
                stack.pop()
                # TODO  此时的 stack[-1] 表示上一个比prev 小的坐标，主要是考虑相等的情况
                # TODO heights[stack[01]] > heights 确保了不会是第一次入栈的 0
                largest = max(largest, (index -  stack[-1]-1) * heights[prev])
            stack.append(index)
        return largest     
# @lc code=end

if __name__ == "__main__":
    print(Solution().largestRectangleArea([2,1,2]))