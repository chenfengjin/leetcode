#
# @lc app=leetcode.cn id=21 lang=python3
#
# [21] 合并两个有序链表
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
# TODO 核心是对重复值得的处理
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        # TODO 这里也ok
        head = ListNode()
        current = head
        while l1 and l2:
            if l1.val > l2.val:
                current.next = l2
                l2 = l2.next
            elif l1.val == l2.val:
                current.next = l1
                l1 = l1.next
                # l2 = l2.next
            else:
                current.next = l1
                l1 = l1.next
            current = current.next
        # TODO 这里写的比较好
        l = l1 if l1 else l2
        # while l:
        #     current.next = l
        #     l = l.next
        # TODO 这里写的比较好
        current.next = l
        return head.next      
# @lc code=end

