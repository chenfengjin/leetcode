#
# @lc app=leetcode id=106 lang=python3
#
# [106] Construct Binary Tree from Inorder and Postorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> TreeNode:
        if not inorder:
            return None
        root = TreeNode(postorder[-1])
        index  = inorder.index(postorder[-1])
        root.left = self.buildTree(inorder =  inorder[:index],postorder=postorder[:index])
        root.right = self.buildTree(inorder = inorder[index+1:],postorder = postorder[index:len(postorder)-1])
        return root
# @lc code=end

