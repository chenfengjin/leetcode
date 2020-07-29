#
# @lc app=leetcode id=92 lang=python3
#
# [92] Reverse Linked List II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:
        current = head
        index = 1
        stack = []
        pivot = None
        while current.next:
            if  m <= index <= n:
                stack.append(current)
            if index == m:
                pivot = current
            if index == n+1 :
                current = pivot
                _next = current
                while stack:
                    current.next = stack.pop()
                    current = current.next
                current.next = _next                
            current = current.next
            index += 1
        return head                 
        
# @lc code=end

