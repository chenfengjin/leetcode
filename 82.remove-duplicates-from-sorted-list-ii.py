#
# @lc app=leetcode id=82 lang=python3
#
# [82] Remove Duplicates from Sorted List II
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if not head:
            return None
        current = head
        left = head
        right = head
        while right.next:
            if left.next.val == right.val:
                right = right.next
            else:
                left.next = right
            right = right.next
        
        
# @lc code=end

