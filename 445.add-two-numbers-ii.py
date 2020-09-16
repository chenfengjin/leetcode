#
# @lc app=leetcode id=445 lang=python3
#
# [445] Add Two Numbers II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        if not l1:
            return l2
        if not l2:
            return l1
        
        length_l1 = 0
        length_l2 = 0
        current = l1
        while current:
            length_l1 += 1
            current = current.next
        current = l2
        while current:
            length_l2 += 1
            current = current.next
        
        if length_l1 < length_l2:
            l1,l2 = l2,l1
            length_l1,length_l2 = length_l2,length_l2
        
        if  length_l1 > length_l2:
            length_diff = length_l1 - length_l2
            head =ListNode(0)
            current = head
            length_diff -= 1
            while length_diff:
                node =ListNode(0)
                current.next = node
                current = current.next
                length_diff -= 1
            current.next = l2
            l2 = head
        next_node,carry = self.helper(l1,l2)
        if not carry:
            return next_node
        node =ListNode(carry)
        node.next = next_node
        return node
    

    def helper(self,l1,l2):
        if not l1 or not l2:
            return None,0
        next_node,carry = self.helper(l1.next,l2.next)
        node =ListNode((l1.val + l2.val + carry) % 10)
        node.next = next_node
        carry = (l1.val + l2.val +carry) // 10
        return node,carry

# @lc code=end

