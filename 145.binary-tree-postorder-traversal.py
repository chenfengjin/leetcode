#
# @lc app=leetcode id=145 lang=python3
#
# [145] Binary Tree Postorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        result = []
        stack = [root]
        while stack:
            node = stack.pop()
            left = node.left
            right = node.right
            node.left = None
            node.right = None
            if left or right:
                stack.append(node)
                if right:
                    stack.append(right)
                if left:
                    stack.append(left)
            else:
                result.append(node.val)
        return result
        
# @lc code=end

