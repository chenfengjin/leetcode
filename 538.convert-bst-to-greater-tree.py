#
# @lc app=leetcode id=538 lang=python3
#
# [538] Convert BST to Greater Tree
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def convertBST(self, root: TreeNode) -> TreeNode:
        if not root:
            return None
        if root.right:
            self.convertBST(root.right)
            root.val += root.right.val
        if root.left:
            root.left.val += root.val
            self.convertBST(root.right)
        return root
        
        
# @lc code=end

