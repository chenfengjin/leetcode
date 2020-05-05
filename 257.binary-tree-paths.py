#
# @lc app=leetcode id=257 lang=python3
#
# [257] Binary Tree Paths
#
import copy
# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.res = []
        self.subpath = []

    def binaryTreePaths(self, root: TreeNode) -> List[str]:
        if not root:
            return []
        if root:
            self.subpath.append(root.val)
        self.helper(root)
        return self.res
        

    def helper(self,root):
        if not root.left and not root.right:
            self.res.append("->".join([str(i) for i in self.subpath]))
            return 
        if root.left:
            self.subpath.append(root.left.val)
            self.helper(root.left)
            self.subpath.pop()
        if root.right:
            self.subpath.append(root.right.val)
            self.helper(root.right)
            self.subpath.pop()
        
        
# @lc code=end

