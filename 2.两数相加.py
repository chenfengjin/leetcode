#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        carry = 0
        l = ListNode(0)
        current = l
        while l1 or l2: #这里用 or 比较好一些
            total =  carry # TODO 这个不错
            total += l1.val if l1 else 0
            total += l2.val if l2 else 0
            # print(current.val)
            # print(current)
            current.next = ListNode(total % 10)
            current = current.next
            carry = total // 10
            l1 = l1.next
            l2 = l2.next
        if carry:
            current.next = ListNode(carry)
        print(l.val)
        print(l.next.val)
        print(l.next.next.val)
        print(l.next.next.next.val)
        return l.next          

# @lc code=end

