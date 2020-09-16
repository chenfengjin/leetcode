#
# @lc app=leetcode id=129 lang=python3
#
# [129] Sum Root to Leaf Numbers
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

#  太丑，想想怎么写好看
class Solution:
    def __init__(self):
        self.total = 0
        self.current = 0
    def sumNumbers(self, root: TreeNode) -> int:
        if not root:
            return 0
        self.current = root.val
        self.helper(root)
        return self.total
        
    def helper(self,root):
        if not root.left and not root.right:
            self.total += self.current
            return 

        self.current *= 10
        if root.left:
            self.current += root.left.val
            self.helper(root.left)
            self.current -= root.left.val

        if root.right:
            self.current += root.right.val
            self.helper(root.right)
            self.current -= root.right.val
        self.current = self.current // 10
        
        
# @lc code=end

