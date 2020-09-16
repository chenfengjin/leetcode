#
# @lc app=leetcode id=98 lang=python3
#
# [98] Validate Binary Search Tree
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.values = []
        self.call_count = 1
        
    def isValidBST(self, root: TreeNode) -> bool:

        self.call_count += 1
        if not root:
            return True
        if not root.left and not  root.right:
            if self.values and self.values[-1] >= root.val:
                return False
            self.values.append(root.val)
            return True
        if root.left:
            if not self.isValidBST(root.left):
                return False
        if self.values and self.values[-1] >= root.val:
            return False
        self.values.append(root.val)
        if root.right:
            if not self.isValidBST(root.right):
                return False
        return True
        
        
# @lc code=end

# if __name__ == "__main__":
#     right = TreeNode(1)
#     root = TreeNode(1)
#     root.right = right
#     print(Solution().isValidBST(root))