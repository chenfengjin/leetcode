#
# @lc app=leetcode id=141 lang=python3
#
# [141] Linked List Cycle
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if not head:
            return False
        left = head
        right = head.next
        while right:
            if left == right:
                return True
            left = left.next
            right = right.next
            if right:
                right = right.next
        return False
        
# @lc code=end

