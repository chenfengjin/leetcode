#
# @lc app=leetcode id=515 lang=python3
#
# [515] Find Largest Value in Each Tree Row
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def largestValues(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        res = []
        current_stack = [root]
        while current_stack:
            layer_values = []
            next_stack = []
            while current_stack:
                node = current_stack.pop()
                layer_values.append(node.val)
                if node.left:
                    next_stack.append(node.left)
                if node.right:
                    next_stack.append(node.right)
            current_stack = next_stack
            res.append(max(layer_values))

        return res
                 
                      
        
       
# @lc code=end

