#
# @lc app=leetcode id=113 lang=python3
#
# [113] Path Sum II
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> List[List[int]]:
        # print(root.val,sum)
        if not root:
            return []
        if not root.left and not root.right:
            if root.val == sum:
                # print(sum)
                return [[root.val]]
            return []
        left =  self.pathSum(root.left,sum-root.val)
        right = self.pathSum(root.right,sum-root.val)
        # 奇丑无比
        [i.insert(0,root.val) for i in left]
        [i.insert(0,root.val) for i in right]
        return left + right
            
    
# @lc code=end

