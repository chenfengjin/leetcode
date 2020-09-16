#
# @lc app=leetcode id=103 lang=python3
#
# [103] Binary Tree Zigzag Level Order Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution: #这个是层序的代码，改吧改吧就好了
    def zigzagLevelOrder(self, root: TreeNode) -> List[List[int]]:
        if not root:
            return []
        # if not root.left and not root.left
        res = []
        current_level_values = []
        current_queue = []
        next_queue = []
        current_queue.append(root)

        index = 0
        while current_queue:
            node = current_queue.pop(index)
            if node.left:
                next_queue.append(node.left)
            if node.right:
                next_queue.append(node.right)
            current_level_values.append(node.val)
            if not current_queue:
                index = 0 if index == -1 else -1

                res.append(current_level_values)
                current_level_values = []
                if next_queue:
                    current_queue = next_queue
                    next_queue = []
        
        return res      
# @lc code=end

