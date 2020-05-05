#
# @lc app=leetcode id=124 lang=python3
#
# [124] Binary Tree Maximum Path Sum
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
        self.max_path_sum = None

    def maxPathSum(self, root: TreeNode) -> int:
        pass
    
    def helper(self,root):
        # if not root:
        #     return 0

        if not root.left and not root.right:
            if self.max_path_sum is None or self.max_path_sum < root.val :
                self.max_path_sum = root.val
            return root.val

        left_subtree_min_value = self.helper(root.left)
        right_subtree_min_value = self.helper(root.right)
        if root.val > self
        


# @lc code=end

