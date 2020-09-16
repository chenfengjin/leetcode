#
# @lc app=leetcode id=101 lang=python3
#
# [101] Symmetric Tree
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        if not root:
            return True
        if not root.left and not root.right:
            return True
        if not root.left or not root.right:
            return False
        
        return self.isMirror(root.left,root.right)

    def isMirror(self,root1:TreeNode,root2:TreeNode):
        if not root1 and not root2:
            return True
        if not root1 or not root2:
            return False
        if not root1.val == root2.val:
            return False
        return self.isMirror(root1.left,root2.right) and self.isMirror(root1.right,root2.left)

        
# @lc code=end

