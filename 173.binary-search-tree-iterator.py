#
# @lc app=leetcode id=173 lang=python3
#
# [173] Binary Search Tree Iterator
#

# @lc code=start
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class BSTIterator:

    def __init__(self, root: TreeNode):
        self.nodes = []
        if root:
            self.nodes.append(root)

    def next(self) -> int:
        """
        @return the next smallest number
        """
        node = self.nodes.pop()
        if not node.left and not node.right:
            return node.val
        left = node.left
        right = node.right
        node.left = None
        node.right = None
        if right:
            self.nodes.append(right)
        if left:
            self.nodes.append(node)
            self.nodes.append(left)
            return self.next()
        return node.val


    def hasNext(self) -> bool:
        """
        @return whether we have a next smallest number
        """
        return len(self.nodes)
        

# Your BSTIterator object will be instantiated and called as such:
# obj = BSTIterator(root)
# param_1 = obj.next()
# param_2 = obj.hasNext()
# @lc code=end


if __name__ == "__main__":
    BSTIterator()