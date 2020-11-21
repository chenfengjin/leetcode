#
# @lc app=leetcode.cn id=102 lang=python3
#
# [102] 二叉树的层序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        current_layer = [root]
        next_layer = []
        result = [[]]
        while current_layer or next_layer:
            if not  current_layer:
                current_layer = next_layer
                next_layer = []
                result.append([])
            if not current_layer:
                return result
            node = current_layer.pop(0)
            if node.left:
                next_layer.append(node.left)
            if node.right:
                next_layer.append(node.right)
            result[-1].append(node.val)
        return result            
# @lc code=end

