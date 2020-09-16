#
# @lc app=leetcode id=230 lang=python3
#
# [230] Kth Smallest Element in a BST
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution: 
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        nodes = [root]
        while nodes:
            node = nodes.pop()
            if not node.left and not node.right:
                k = k - 1
                if k == 0:
                    return node.val
                continue
            left = node.left
            right = node.right 
            node.left = None
            node.right = None
            if right:
                nodes.append(right)
            nodes.append(node)
            if left:
                nodes.append(left)
        
        
# @lc code=end

