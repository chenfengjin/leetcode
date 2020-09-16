#
# @lc app=leetcode.cn id=19 lang=python3
#
# [19] 删除链表的倒数第N个节点
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
# 需要考虑删第一个的场景
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        l = ListNode(0)
        l.next = head
        right = l
        left = l
        while n:
            n = n-1
            right = right.next
        while right.next:
            right = right.next
            left = left.next
        left.next = left.next.next
        return l.next
        
# @lc code=end
