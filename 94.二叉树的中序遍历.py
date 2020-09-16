#
# @lc app=leetcode.cn id=94 lang=python3
#
# [94] 二叉树的中序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        s = []
        result = []
        while root or len(s):
            while root:
                s.append(root)
                root = root.left
            root = s[-1]
            s.pop()
            result.append(root.val)
            root = root.right
        return result
# 很优雅的写法
# @lc code=end

