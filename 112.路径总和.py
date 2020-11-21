#
# @lc app=leetcode.cn id=112 lang=python3
#
# [112] 路径总和
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool: #不想纠结时间了，直接超的路径总和II
        if not root:
            return False
        result= []
        current = []
        self.dfs(result,current,root,sum)
        return not len(result) == 0
    def dfs(self,result,current,root,sum):
        
        current.append(root.val)
        sum -= root.val 
        if not root.left and not root.right:
            if sum == 0:
                result.append(current[:])
            # 中途 return 记得恢复现场
            # 然而还是错的，因为可能有负数
            # return 
        if root.left:
            self.dfs(result,current,root.left,sum)
        if root.right:
            self.dfs(result,current,root.right,sum)
        current.pop()
# @lc code=end

