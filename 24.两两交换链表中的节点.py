#
# @lc app=leetcode.cn id=24 lang=python3
#
# [24] 两两交换链表中的节点
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        l = ListNode(0)
        l.next = head #TODO 基本操作了
        current = l 
        while current.next and current.next.next:
            t = current.next.next.next
            current.next.next.next = current.next
            current.next = current.next.next
            current.next.next.next = t #TODO 需要注意指针是变化的，最好画一下
            current = current.next 
            current = current.next
        return l.next
# @lc code=end

