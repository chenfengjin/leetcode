#
# @lc app=leetcode.cn id=203 lang=python3
#
# [203] 移除链表元素
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        h = ListNode(0)
        h.next = head            
        current = h
        while current.next:
            if current.next.val == val:
                current.next = current.next.next
            else:
                current = current.next
        return h.next   
        
# @lc code=end

