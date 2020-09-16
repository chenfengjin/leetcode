#
# @lc app=leetcode id=117 lang=python3
#
# [117] Populating Next Right Pointers in Each Node II
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
class Solution:
    def connect(self, root: 'Node') -> 'Node':
        if not root:
            return 
        res = []
        current_level_values = []
        current_queue = []
        next_queue = []
        current_queue.append(root)

        while current_queue:
            node = current_queue.pop(0)
            if current_queue:
                node.next = current_queue[0]
            # else:
                # node.next = None
            if node.left:
                next_queue.append(node.left)
            if node.right:
                next_queue.append(node.right)
            current_level_values.append(node.val)
            if not current_queue:
                res.append(current_level_values)
                current_level_values = []
                if next_queue:
                    current_queue = next_queue
                    next_queue = []
        
        return  root                
# @lc code=end

