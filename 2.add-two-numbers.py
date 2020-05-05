#
# @lc app=leetcode id=2 lang=python3
#
# [2] Add Two Numbers
#

# @lc code=start
# Definition for singly-linked list.
<<<<<<< HEAD
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

        
=======
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        if l1.val == 0 and not l1.next:
            return l2
        if l2.val == 0 and not l2.next:
            return l1
        carry = 0
        l3 = ListNode((l1.val+l2.val)%10)
        carry = (l1.val + l2.val)//10
        current = l3
        l1 = l1.next
        l2 = l2.next
        while l1 and l2:
            node = ListNode((l1.val + l2.val + carry)%10)
            carry = (l1.val + l2.val + carry) // 10
            l1=l1.next
            l2 = l2.next   
            current.next = node 
            current = current.next
        remain = l1 if l1 else (l2 if l2 else None)

        while remain and carry:
            node = ListNode((remain.val + carry)%10)
            carry = (remain.val + carry) // 10
            remain=remain.next
            current.next = node 
            current = current.next
        if carry:
            current.next = ListNode(carry)
        if remain:
            current.next = remain   

        return l3

            

        
        
>>>>>>> RFC:change mac
# @lc code=end

