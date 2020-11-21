#
# @lc app=leetcode.cn id=42 lang=python3
#
# [42] 接雨水
#

# @lc code=start
class Solution:
    def trap(self, height: List[int]) -> int:
        # 单调栈
        height = [0] + height + [0]
        stack = []
        total = 0
        for index,h in enumerate(height):
            while stack  and height(stack[-1]) < height:
                current = stack[-1]
                depth = min(height[current+1],height[index-1])
                total += depth * (index-current) 
            stack.append(index)
        return total

# @lc code=end

