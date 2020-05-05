#
# @lc app=leetcode id=61 lang=python3
#
# [61] Rotate List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution: #背诵一下这个题目把，主要是考虑边界条件
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        if not head:
            return head
        length = 1
        current = head

        while current.next:
            current = current.next
            length += 1

        current.next = head

        pivot = head
        k  =  k % length
        while length > k + 1:
            pivot = pivot.next
            length = length - 1 

        head = pivot.next
        pivot.next = None

        return head


# @lc code=end

