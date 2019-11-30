#
# @lc app=leetcode id=2 lang=python3
#
# [2] Add Two Numbers
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        result = ListNode(0)
        carry = 0
        r=result
        while l1 and l2:
            current = ListNode(0)
            current.val = (l1.val+l2.val+carry) // 10
            carry  = (l1.val+l2.val + carry )%10
            l1 = l1.next
            l2 = l2.next
            current.next = ListNode(0)
            current = current.next
            r = r.next
            r=current
            
        if l1:
            current = l1
        if l2:
            current = l2
        return result

        
# @lc code=end

