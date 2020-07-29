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
            return head

        head1 = ListNode(None)
        head1.next = head
        head = head1

        current = head

        
        skip = False
        while current.next and current.next.next:
            if current.next.val == current.next.next.val:
                skip = True
                current.next = current.next.next
            else:
                if skip:
                    skip = False
                    current.next = current.next.next
                else:
                    current = current.next
        return head.next
        
        
# @lc code=end

