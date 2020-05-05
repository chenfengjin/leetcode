#
# @lc app=leetcode id=19 lang=python3
#
# [19] Remove Nth Node From End of List
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        if not head:
            return None
        if not head.next and n == 1:
            return None
        # a[0]=1
        left  = head
        right = head
        previous = head
        for _ in range(n):
            right = right.next
        while right.next:
            right = right.next
            previous = left
            left = left.next
        if left:
            try:
                previous.next = left.next
            except:
                pass
        return head

        
        
# @lc code=end

