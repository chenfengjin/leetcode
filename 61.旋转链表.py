#
# @lc app=leetcode.cn id=61 lang=python3
#
# [61] 旋转链表
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
# 边界条件是关键
class Solution:
    def rotateRight(self, head: ListNode, k: int) -> ListNode:
        l = head
        count = 0
        while l:
            count += 1
            l = l.next
        if count == 0: #TODO bug point
            return None
        k %= count
        if not k: # TODO bug point
            return head
        right = head
        while k:
            right  = right.next
            k = k - 1
        left = head
        while right.next:
            left = left.next
            right = right.next
        h = left.next
        left.next = None #TODO bugpoint
        right.next = head
        return h


# @lc code=end

# 1->2->3->4->5 