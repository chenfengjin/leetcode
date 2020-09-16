#
# @lc app=leetcode.cn id=98 lang=python3
#
# [98] 验证二叉搜索树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        # if not root:
        #     return True

        # if root.left and not root.val > root.left.val:
        #     return False

        # if root.right and not root.val < root.right.val:
        #     return False
        # return self.isValidBST(root.left) and self.isValidBST(root.right)
        s = []
        current = float('-inf')
        while root or len(s):
            while root:
                s.append(root)
                root = root.left
            root = s[-1]
            s.pop()
            if current >= root.val:
                return False
            current = root.val
            root = root.right  
        return True        
# @lc code=end

