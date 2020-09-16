#
# @lc app=leetcode id=94 lang=python3
#
# [94] Binary Tree Inorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution: 
    # def __init__(self):
    #     self.elements = []

    # def inorderTraversal(self, root: TreeNode) -> List[int]:
    #     self.helper(root)
    #     return self.elements

    # def helper(self,root):
    #     if not root:
    #         return
    #     self.helper(root.left)
    #     self.elements.append(root.val)
    #     self.helper(root.right)
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        if not root:
            return []
        
        nodes_to_search = [root]
        elements = []

        while nodes_to_search:
            node_to_search = nodes_to_search.pop()
            if node_to_search.left:
                node_to_search.left = None
                left = node_to_search.left
                nodes_to_search.append(node_to_search)
                nodes_to_search.append(left)
                continue
            elements.append(node_to_search.val)

            if node_to_search.right:
                nodes_to_search.append(node_to_search.right)
# @lc code=end

