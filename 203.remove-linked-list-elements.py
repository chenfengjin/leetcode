#
# @lc app=leetcode id=203 lang=python3
#
# [203] Remove Linked List Elements
#

# @lc code=start
# Definition for singly-linked list.
class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        head = ListNode(0,head)
        current = head
        while current.next:
            if current.next.val == val:
                current.next = current.next.next
            else:
                current = current.next
        return head.next
        
# @lc code=end

